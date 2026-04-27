// Package interactive handles Tabstack agent.automate's mid-stream
// "interactive:form_data:request" events. The package defines neutral
// FormDataRequest/FormDataField types (decoded from the SDK's untyped
// event payload), a Prompter abstraction (TTY or file-source), and a
// thin Submitter wrapper around the SDK's Agent.AutomateInput call.
package interactive

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	tabstack "github.com/stainless-sdks/tabstack-go"
	"github.com/stainless-sdks/tabstack-go/packages/param"
	"golang.org/x/term"
)

// FormDataRequest is the decoded payload of an
// "interactive:form_data:request" event. The struct tags match the API's
// JSON shape so DecodeFormDataRequest can JSON-roundtrip the SDK's
// untyped event Data into this type.
type FormDataRequest struct {
	RequestID       string          `json:"requestId"`
	IterationID     string          `json:"iterationId"`
	PageTitle       string          `json:"pageTitle"`
	PageURL         string          `json:"pageUrl"`
	FormDescription string          `json:"formDescription"`
	Fields          []FormDataField `json:"fields"`
}

// FormDataField is a single field the agent needs data for.
type FormDataField struct {
	Ref          string   `json:"ref"`
	Label        string   `json:"label"`
	FieldType    string   `json:"fieldType"`
	Required     bool     `json:"required"`
	CurrentValue string   `json:"currentValue"`
	Description  string   `json:"description"`
	Options      []string `json:"options"`
}

// FieldValue is a single {ref, value} answer for a form field.
type FieldValue struct {
	Ref   string
	Value string
}

// DecodeFormDataRequest converts an SDK event Data payload (typed as `any`,
// usually a map[string]any from JSON unmarshalling) into a FormDataRequest.
func DecodeFormDataRequest(data any) (FormDataRequest, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return FormDataRequest{}, fmt.Errorf("marshal event data: %w", err)
	}
	var req FormDataRequest
	if err := json.Unmarshal(b, &req); err != nil {
		return FormDataRequest{}, fmt.Errorf("unmarshal event data: %w", err)
	}
	return req, nil
}

// Prompter resolves form-data requests into field values.
type Prompter interface {
	// Prompt returns either a list of field values or a cancellation. An
	// error means the prompt itself failed (e.g. read error); the caller
	// must not submit anything in that case.
	Prompt(req FormDataRequest) ([]FieldValue, bool, error)
}

// Submitter posts a response to /automate/{requestId}/input.
type Submitter interface {
	Submit(ctx context.Context, requestID string, fields []FieldValue, cancelled bool) error
}

// HandleRequest runs the prompter and forwards its result to the submitter.
// On prompt error the submitter is NOT called.
func HandleRequest(ctx context.Context, s Submitter, p Prompter, req FormDataRequest) error {
	fields, cancelled, err := p.Prompt(req)
	if err != nil {
		return fmt.Errorf("prompt for request %s: %w", req.RequestID, err)
	}
	if cancelled {
		fields = nil
	}
	return s.Submit(ctx, req.RequestID, fields, cancelled)
}

// --- TTY prompter ---

// TTYPrompter prompts on stderr and reads from stdin.
type TTYPrompter struct {
	in  io.Reader
	out io.Writer
}

// NewTTYPrompter constructs a prompter that reads from os.Stdin and writes
// prompts to os.Stderr (so stdout stays reserved for SSE events).
func NewTTYPrompter() *TTYPrompter {
	return &TTYPrompter{in: os.Stdin, out: os.Stderr}
}

func (t *TTYPrompter) Prompt(req FormDataRequest) ([]FieldValue, bool, error) {
	_, _ = fmt.Fprintf(t.out, "\n=== Tabstack form-data request %s ===\n", req.RequestID)
	_, _ = fmt.Fprintf(t.out, "Page: %s (%s)\n", req.PageTitle, req.PageURL)
	if req.FormDescription != "" {
		_, _ = fmt.Fprintf(t.out, "Form: %s\n", req.FormDescription)
	}
	_, _ = fmt.Fprintln(t.out, `Type "/cancel" at any prompt to cancel the request.`)

	reader := bufio.NewReader(t.in)
	values := make([]FieldValue, 0, len(req.Fields))
	for _, f := range req.Fields {
		label := f.Label
		if label == "" {
			label = f.Ref
		}
		marker := "(required)"
		if !f.Required {
			marker = "(optional)"
		}
		extra := ""
		if f.Description != "" {
			extra = " — " + f.Description
		}
		if len(f.Options) > 0 {
			extra += " options: " + strings.Join(f.Options, ", ")
		}
		_, _ = fmt.Fprintf(t.out, "  %s %s [%s]%s\n  > ", label, marker, f.FieldType, extra)

		v, eof, err := t.readField(reader, f)
		if err != nil {
			return nil, false, fmt.Errorf("read input: %w", err)
		}
		if eof {
			return nil, true, nil // EOF — treat as cancel
		}
		if v == "/cancel" {
			return nil, true, nil
		}
		values = append(values, FieldValue{Ref: f.Ref, Value: v})
	}
	return values, false, nil
}

// readField reads one line of input. For fields whose type indicates a secret
// (currently "password"), the line is read without echo via term.ReadPassword
// when t.in is an *os.File attached to a TTY. For non-TTY inputs (pipes,
// redirected files, tests) and non-secret fields, falls back to a plain
// buffered line read.
func (t *TTYPrompter) readField(reader *bufio.Reader, f FormDataField) (value string, eof bool, err error) {
	if f.FieldType == "password" {
		if file, ok := t.in.(*os.File); ok && term.IsTerminal(int(file.Fd())) {
			pw, perr := term.ReadPassword(int(file.Fd()))
			if perr != nil {
				return "", false, perr
			}
			// User pressed Enter but it wasn't echoed; emit a newline so the
			// next prompt line doesn't run on.
			_, _ = fmt.Fprintln(t.out)
			return string(pw), false, nil
		}
		// Non-TTY: read plainly. Echo behavior is the test/pipe's concern.
	}
	line, rerr := reader.ReadString('\n')
	if rerr != nil && rerr != io.EOF {
		return "", false, rerr
	}
	if len(line) == 0 && rerr == io.EOF {
		return "", true, nil
	}
	return strings.TrimRight(line, "\r\n"), false, nil
}

// --- File prompter ---

// FilePrompter looks up answers in a JSON file keyed by request ID.
// File shape:
//
//	{
//	  "<requestId>": { "<fieldRef>": "<value>", ... },
//	  ...
//	}
//
// A missing request ID is treated as a cancel.
type FilePrompter struct {
	answers map[string]map[string]string
}

// NewFilePrompter reads a JSON answers file from path.
func NewFilePrompter(path string) (*FilePrompter, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read input file: %w", err)
	}
	var m map[string]map[string]string
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, fmt.Errorf("parse input file: %w", err)
	}
	return &FilePrompter{answers: m}, nil
}

func (f *FilePrompter) Prompt(req FormDataRequest) ([]FieldValue, bool, error) {
	answers, ok := f.answers[req.RequestID]
	if !ok {
		return nil, true, nil
	}
	values := make([]FieldValue, 0, len(req.Fields))
	for _, fld := range req.Fields {
		v, ok := answers[fld.Ref]
		if !ok {
			// Missing answer for a required field is treated as cancel —
			// submitting a partial response would 4xx from the API. Optional
			// fields are silently skipped.
			if fld.Required {
				return nil, true, nil
			}
			continue
		}
		values = append(values, FieldValue{Ref: fld.Ref, Value: v})
	}
	return values, false, nil
}

// --- SDK submitter ---

// SDKSubmitter implements Submitter against the live Tabstack client.
type SDKSubmitter struct {
	Client *tabstack.Client
}

// Submit posts form-data input (or a cancellation) to the Tabstack API.
func (s SDKSubmitter) Submit(ctx context.Context, requestID string, fields []FieldValue, cancelled bool) error {
	body := tabstack.AgentAutomateInputParams{}
	if cancelled {
		body.Cancelled = param.NewOpt(true)
	} else {
		body.Fields = make([]tabstack.AgentAutomateInputParamsField, 0, len(fields))
		for _, f := range fields {
			body.Fields = append(body.Fields, tabstack.AgentAutomateInputParamsField{
				Ref:   param.NewOpt(f.Ref),
				Value: param.NewOpt(f.Value),
			})
		}
	}
	_, err := s.Client.Agent.AutomateInput(ctx, requestID, body)
	return err
}
