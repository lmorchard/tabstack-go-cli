package interactive

import (
	"bytes"
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func makeRequest(reqID string, fieldRefs ...string) FormDataRequest {
	fields := make([]FormDataField, 0, len(fieldRefs))
	for _, ref := range fieldRefs {
		fields = append(fields, FormDataField{
			Ref:       ref,
			Label:     "Field " + ref,
			FieldType: "text",
			Required:  true,
		})
	}
	return FormDataRequest{
		RequestID: reqID,
		Fields:    fields,
	}
}

func TestDecodeFormDataRequest_RoundTrip(t *testing.T) {
	raw := map[string]any{
		"requestId":       "req-1",
		"pageTitle":       "Sign in",
		"pageUrl":         "https://example.com/login",
		"formDescription": "Login form",
		"fields": []any{
			map[string]any{
				"ref":       "E1",
				"label":     "Email",
				"fieldType": "email",
				"required":  true,
			},
			map[string]any{
				"ref":       "E2",
				"label":     "Password",
				"fieldType": "password",
				"required":  true,
			},
		},
	}
	got, err := DecodeFormDataRequest(raw)
	if err != nil {
		t.Fatalf("DecodeFormDataRequest: %v", err)
	}
	if got.RequestID != "req-1" {
		t.Errorf("RequestID: %q", got.RequestID)
	}
	if got.PageTitle != "Sign in" || got.PageURL != "https://example.com/login" {
		t.Errorf("page: %q / %q", got.PageTitle, got.PageURL)
	}
	if len(got.Fields) != 2 {
		t.Fatalf("fields: %d", len(got.Fields))
	}
	if got.Fields[0].Ref != "E1" || got.Fields[0].FieldType != "email" {
		t.Errorf("field 0: %+v", got.Fields[0])
	}
}

func TestFilePrompter_AnswersByRequestID(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "answers.json")
	body := []byte(`{
		"req-1": {"E1": "alice", "E2": "alice@example.com"},
		"req-2": {"E1": "bob"}
	}`)
	if err := os.WriteFile(path, body, 0o600); err != nil {
		t.Fatal(err)
	}

	p, err := NewFilePrompter(path)
	if err != nil {
		t.Fatalf("NewFilePrompter: %v", err)
	}
	got, cancelled, err := p.Prompt(makeRequest("req-1", "E1", "E2"))
	if err != nil {
		t.Fatalf("Prompt: %v", err)
	}
	if cancelled {
		t.Fatal("unexpected cancel")
	}
	if len(got) != 2 {
		t.Fatalf("got %d fields, want 2", len(got))
	}
	values := map[string]string{}
	for _, f := range got {
		values[f.Ref] = f.Value
	}
	if values["E1"] != "alice" || values["E2"] != "alice@example.com" {
		t.Errorf("values: %+v", values)
	}
}

func TestFilePrompter_MissingRequestIDCancels(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "answers.json")
	if err := os.WriteFile(path, []byte(`{}`), 0o600); err != nil {
		t.Fatal(err)
	}
	p, err := NewFilePrompter(path)
	if err != nil {
		t.Fatal(err)
	}
	_, cancelled, err := p.Prompt(makeRequest("req-missing", "E1"))
	if err != nil {
		t.Fatalf("Prompt: %v", err)
	}
	if !cancelled {
		t.Error("expected cancel for missing request id")
	}
}

// FilePrompter must cancel when a required field is missing from the answers
// file rather than submitting a partial response (which would 4xx).
func TestFilePrompter_MissingRequiredFieldCancels(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "answers.json")
	// Provide an answer for E1 but not for E2 (which is required).
	if err := os.WriteFile(path, []byte(`{"req-1": {"E1": "alice"}}`), 0o600); err != nil {
		t.Fatal(err)
	}
	p, err := NewFilePrompter(path)
	if err != nil {
		t.Fatal(err)
	}
	_, cancelled, err := p.Prompt(makeRequest("req-1", "E1", "E2"))
	if err != nil {
		t.Fatalf("Prompt: %v", err)
	}
	if !cancelled {
		t.Error("expected cancel when a required field has no answer")
	}
}

// TTYPrompter falls back to a plain line read for password fields when stdin
// isn't a real TTY (e.g., in tests, or under a pipe). The no-echo path itself
// requires a real terminal and isn't exercised here.
func TestTTYPrompter_PasswordNonTTYFallback(t *testing.T) {
	in := strings.NewReader("secret\n")
	var out bytes.Buffer
	p := &TTYPrompter{in: in, out: &out}
	req := FormDataRequest{
		RequestID: "req-pw",
		Fields: []FormDataField{
			{Ref: "E1", Label: "Password", FieldType: "password", Required: true},
		},
	}
	got, cancelled, err := p.Prompt(req)
	if err != nil {
		t.Fatalf("Prompt: %v", err)
	}
	if cancelled {
		t.Fatal("unexpected cancel")
	}
	if len(got) != 1 || got[0].Ref != "E1" || got[0].Value != "secret" {
		t.Errorf("got %+v, want E1=secret", got)
	}
}

// Optional fields with no answer should be silently skipped, not cancel.
func TestFilePrompter_MissingOptionalFieldSkipped(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "answers.json")
	if err := os.WriteFile(path, []byte(`{"req-1": {"E1": "alice"}}`), 0o600); err != nil {
		t.Fatal(err)
	}
	p, err := NewFilePrompter(path)
	if err != nil {
		t.Fatal(err)
	}
	req := FormDataRequest{
		RequestID: "req-1",
		Fields: []FormDataField{
			{Ref: "E1", FieldType: "text", Required: true},
			{Ref: "E2", FieldType: "text", Required: false},
		},
	}
	got, cancelled, err := p.Prompt(req)
	if err != nil {
		t.Fatalf("Prompt: %v", err)
	}
	if cancelled {
		t.Error("did not expect cancel — optional field should be skipped silently")
	}
	if len(got) != 1 || got[0].Ref != "E1" || got[0].Value != "alice" {
		t.Errorf("got %+v, want a single E1=alice", got)
	}
}

type fakeSubmitter struct {
	calls []submitCall
	err   error
}

type submitCall struct {
	requestID string
	fields    []FieldValue
	cancelled bool
}

func (f *fakeSubmitter) Submit(_ context.Context, requestID string, fields []FieldValue, cancelled bool) error {
	f.calls = append(f.calls, submitCall{requestID, fields, cancelled})
	return f.err
}

func TestHandleRequest_PromptThenSubmit(t *testing.T) {
	prompter := stubPrompter{
		fields:    []FieldValue{{Ref: "E1", Value: "hello"}},
		cancelled: false,
	}
	sub := &fakeSubmitter{}
	err := HandleRequest(context.Background(), sub, prompter, makeRequest("req-x", "E1"))
	if err != nil {
		t.Fatalf("HandleRequest: %v", err)
	}
	if len(sub.calls) != 1 {
		t.Fatalf("submits: %d", len(sub.calls))
	}
	if sub.calls[0].requestID != "req-x" {
		t.Errorf("requestID: %s", sub.calls[0].requestID)
	}
	if sub.calls[0].cancelled {
		t.Error("expected non-cancelled submit")
	}
}

func TestHandleRequest_CancelledPath(t *testing.T) {
	prompter := stubPrompter{cancelled: true}
	sub := &fakeSubmitter{}
	err := HandleRequest(context.Background(), sub, prompter, makeRequest("req-y", "E1"))
	if err != nil {
		t.Fatalf("HandleRequest: %v", err)
	}
	if !sub.calls[0].cancelled {
		t.Error("expected cancelled submit")
	}
	if len(sub.calls[0].fields) != 0 {
		t.Errorf("expected no fields on cancel, got %v", sub.calls[0].fields)
	}
}

func TestHandleRequest_PromptError(t *testing.T) {
	prompter := stubPrompter{err: errors.New("prompt blew up")}
	sub := &fakeSubmitter{}
	err := HandleRequest(context.Background(), sub, prompter, makeRequest("req-z", "E1"))
	if err == nil {
		t.Fatal("expected error")
	}
	if len(sub.calls) != 0 {
		t.Error("submitter should not be called on prompt error")
	}
}

type stubPrompter struct {
	fields    []FieldValue
	cancelled bool
	err       error
}

func (s stubPrompter) Prompt(_ FormDataRequest) ([]FieldValue, bool, error) {
	return s.fields, s.cancelled, s.err
}
