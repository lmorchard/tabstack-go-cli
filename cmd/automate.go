package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/lmorchard/tabstack-go-cli/internal/client"
	"github.com/lmorchard/tabstack-go-cli/internal/interactive"
	"github.com/lmorchard/tabstack-go-cli/internal/schema"
	"github.com/lmorchard/tabstack-go-cli/internal/sse"
	"github.com/spf13/cobra"
	tabstack "github.com/stainless-sdks/tabstack-go"
	"github.com/stainless-sdks/tabstack-go/packages/param"
)

var (
	automateURL         string
	automateGuardrails  string
	automateDataPath    string
	automateInteractive bool
	automateInputFrom   string
	automateGeo         string
	automateMaxIter     int64
	automateMaxValid    int64
	automateOutput      string
	automateColor       string
)

var automateCmd = &cobra.Command{
	Use:   "automate <task>",
	Short: "Stream an AI browser-automation run",
	Args:  cobra.ExactArgs(1),
	RunE:  runAutomate,
}

func runAutomate(_ *cobra.Command, args []string) error {
	if err := validateEnum("output", automateOutput, validStreamOutputs); err != nil {
		return err
	}
	if err := validateEnum("color", automateColor, validColorModes); err != nil {
		return err
	}
	ctx := context.Background()

	c, err := client.New(GetConfig())
	if err != nil {
		return err
	}

	body := tabstack.AgentAutomateParams{
		Task: args[0],
	}
	if automateURL != "" {
		body.URL = param.NewOpt(automateURL)
	}
	if automateGuardrails != "" {
		body.Guardrails = param.NewOpt(automateGuardrails)
	}
	if automateInteractive {
		body.Interactive = param.NewOpt(true)
	}
	if automateMaxIter > 0 {
		body.MaxIterations = param.NewOpt(automateMaxIter)
	}
	if automateMaxValid > 0 {
		body.MaxValidationAttempts = param.NewOpt(automateMaxValid)
	}
	if automateGeo != "" {
		body.GeoTarget = tabstack.AgentAutomateParamsGeoTarget{
			Country: param.NewOpt(automateGeo),
		}
	}
	if automateDataPath != "" {
		sch, err := schema.Load(automateDataPath)
		if err != nil {
			return err
		}
		body.Data = sch
	}

	prompter, err := selectPrompter(automateInteractive, automateInputFrom)
	if err != nil {
		return err
	}
	submitter := interactive.SDKSubmitter{Client: c}

	stream := c.Agent.AutomateStreaming(ctx, body)
	defer func() { _ = stream.Close() }()

	enc := json.NewEncoder(os.Stdout)
	started := time.Now()
	color := resolveStreamColor(automateColor)
	for stream.Next() {
		ev := stream.Current()
		switch automateOutput {
		case "pretty":
			if err := sse.PrettyAutomate(os.Stdout, ev, started, color); err != nil {
				return fmt.Errorf("render event: %w", err)
			}
		default:
			if err := enc.Encode(ev); err != nil {
				return fmt.Errorf("encode event: %w", err)
			}
		}
		if prompter == nil {
			continue
		}
		if ev.Event != "interactive:form_data:request" {
			continue
		}
		req, err := interactive.DecodeFormDataRequest(ev.Data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "interactive: decode form-data event: %v\n", err)
			continue
		}
		if err := interactive.HandleRequest(ctx, submitter, prompter, req); err != nil {
			fmt.Fprintf(os.Stderr, "interactive: %v\n", err)
		}
	}
	if err := stream.Err(); err != nil {
		return fmt.Errorf("automate stream: %w", err)
	}
	return nil
}

// selectPrompter returns nil when no prompter is needed (interactive mode off
// or no callback expected). It only returns a prompter when --interactive was
// set; otherwise the API will not emit form-data events and there's nothing
// to handle.
func selectPrompter(interactiveMode bool, inputFrom string) (interactive.Prompter, error) {
	if !interactiveMode {
		return nil, nil
	}
	if inputFrom != "" {
		return interactive.NewFilePrompter(inputFrom)
	}
	if isTerminal(os.Stdin) {
		return interactive.NewTTYPrompter(), nil
	}
	// Interactive mode requested but no TTY and no --input-from. Fail-safe by
	// auto-cancelling any callbacks rather than blocking forever on stdin.
	fmt.Fprintln(os.Stderr, "warning: --interactive set but stdin is not a TTY and --input-from is empty; form-data requests will be auto-cancelled")
	return autoCancelPrompter{}, nil
}

type autoCancelPrompter struct{}

func (autoCancelPrompter) Prompt(_ interactive.FormDataRequest) ([]interactive.FieldValue, bool, error) {
	return nil, true, nil
}

func isTerminal(f *os.File) bool {
	fi, err := f.Stat()
	if err != nil {
		return false
	}
	return (fi.Mode() & os.ModeCharDevice) != 0
}

// --- automate input subcommand ---

var (
	automateInputValuesPath string
	automateInputCancel     bool
)

var automateInputCmd = &cobra.Command{
	Use:   "input <request-id>",
	Short: "Submit a response to an interactive form-data request",
	Long: `Reply to (or cancel) an in-flight 'automate' form-data request, identified
by the requestId emitted in an interactive:form_data:request SSE event.

Provide either --values <file> (JSON: {"E1":"alice","E2":"bob"}) or --cancel.`,
	Args: cobra.ExactArgs(1),
	RunE: runAutomateInput,
}

func runAutomateInput(_ *cobra.Command, args []string) error {
	c, err := client.New(GetConfig())
	if err != nil {
		return err
	}

	body := tabstack.AgentAutomateInputParams{}
	switch {
	case automateInputCancel && automateInputValuesPath != "":
		return fmt.Errorf("--cancel and --values are mutually exclusive")
	case automateInputCancel:
		body.Cancelled = param.NewOpt(true)
	case automateInputValuesPath != "":
		raw, err := os.ReadFile(automateInputValuesPath)
		if err != nil {
			return fmt.Errorf("read values: %w", err)
		}
		var values map[string]string
		if err := json.Unmarshal(raw, &values); err != nil {
			return fmt.Errorf("parse values as JSON object: %w", err)
		}
		body.Fields = make([]tabstack.AgentAutomateInputParamsField, 0, len(values))
		for ref, v := range values {
			body.Fields = append(body.Fields, tabstack.AgentAutomateInputParamsField{
				Ref:   param.NewOpt(ref),
				Value: param.NewOpt(v),
			})
		}
	default:
		return fmt.Errorf("provide --values <file> or --cancel")
	}

	_, err = c.Agent.AutomateInput(context.Background(), args[0], body)
	if err != nil {
		return fmt.Errorf("submit input: %w", err)
	}
	fmt.Fprintln(os.Stderr, "ok")
	return nil
}

func init() {
	automateCmd.Flags().StringVar(&automateURL, "url", "", "starting URL for the task")
	automateCmd.Flags().StringVar(&automateGuardrails, "guardrails", "", "safety constraints for execution")
	automateCmd.Flags().StringVar(&automateDataPath, "data", "", `path to JSON file passed as 'data' context for the task (use "-" for stdin)`)
	automateCmd.Flags().BoolVar(&automateInteractive, "interactive", false, "enable interactive form-data callbacks (TTY prompts; combine with --input-from for non-TTY)")
	automateCmd.Flags().StringVar(&automateInputFrom, "input-from", "", `path to JSON file with answers for interactive form-data requests`)
	automateCmd.Flags().StringVar(&automateGeo, "geo", "", "ISO 3166-1 alpha-2 country code")
	automateCmd.Flags().Int64Var(&automateMaxIter, "max-iterations", 0, "max task iterations (0 = SDK default)")
	automateCmd.Flags().Int64Var(&automateMaxValid, "max-validation-attempts", 0, "max validation attempts (0 = SDK default)")
	automateCmd.Flags().StringVar(&automateOutput, "output", "json", "stream output format: json (one event per line) or pretty (human-readable)")
	automateCmd.Flags().StringVar(&automateColor, "color", "auto", "color in pretty output: auto (TTY only, respects NO_COLOR), always, or never")

	automateInputCmd.Flags().StringVar(&automateInputValuesPath, "values", "", `path to JSON file mapping field ref -> value`)
	automateInputCmd.Flags().BoolVar(&automateInputCancel, "cancel", false, "cancel the request instead of providing values")

	automateCmd.AddCommand(automateInputCmd)
	rootCmd.AddCommand(automateCmd)
}
