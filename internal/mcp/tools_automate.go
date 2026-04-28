package mcp

import (
	"context"
	"fmt"
	"os"

	sdk "github.com/modelcontextprotocol/go-sdk/mcp"
	tabstack "github.com/stainless-sdks/tabstack-go"
	"github.com/stainless-sdks/tabstack-go/packages/param"
)

// AutomateInput is the input for tabstack_automate.
type AutomateInput struct {
	Task                  string         `json:"task" jsonschema:"natural-language description of the browser-automation task"`
	URL                   string         `json:"url,omitempty" jsonschema:"starting URL for the task"`
	Guardrails            string         `json:"guardrails,omitempty" jsonschema:"safety constraints for execution"`
	Data                  map[string]any `json:"data,omitempty" jsonschema:"JSON context the agent can use during the task (e.g. for form-filling)"`
	Geo                   string         `json:"geo,omitempty" jsonschema:"ISO 3166-1 alpha-2 country code"`
	MaxIterations         int64          `json:"max_iterations,omitempty" jsonschema:"max task iterations (0 = SDK default)"`
	MaxValidationAttempts int64          `json:"max_validation_attempts,omitempty" jsonschema:"max validation attempts (0 = SDK default)"`
}

func registerAutomate(s *sdk.Server, c *tabstack.Client) {
	sdk.AddTool(s, &sdk.Tool{
		Name: "tabstack_automate",
		Description: "Run an autonomous AI browser-automation task described in natural language. " +
			"The agent navigates web pages, interacts with elements, and returns a final answer or report. " +
			"Tasks requiring user-typed form data (passwords, etc.) are not supported in this MCP version — " +
			"use the `tabstack` CLI directly for those.",
	}, func(ctx context.Context, req *sdk.CallToolRequest, in *AutomateInput) (*sdk.CallToolResult, any, error) {
		body := tabstack.AgentAutomateParams{Task: in.Task}
		// Always non-interactive in MCP. Form-data callbacks are auto-declined
		// below since we have no way to elicit user input mid-tool-call (yet).
		body.Interactive = param.NewOpt(false)
		if in.URL != "" {
			body.URL = param.NewOpt(in.URL)
		}
		if in.Guardrails != "" {
			body.Guardrails = param.NewOpt(in.Guardrails)
		}
		if in.Data != nil {
			body.Data = in.Data
		}
		if in.Geo != "" {
			body.GeoTarget = tabstack.AgentAutomateParamsGeoTarget{Country: param.NewOpt(in.Geo)}
		}
		if in.MaxIterations > 0 {
			body.MaxIterations = param.NewOpt(in.MaxIterations)
		}
		if in.MaxValidationAttempts > 0 {
			body.MaxValidationAttempts = param.NewOpt(in.MaxValidationAttempts)
		}

		emitter := newProgressEmitter(ctx, req)
		stream := c.Agent.AutomateStreaming(ctx, body)
		defer func() { _ = stream.Close() }()

		var finalAnswer string
		for stream.Next() {
			if err := ctx.Err(); err != nil {
				return toolError(err), nil, nil
			}
			ev := stream.Current()
			if msg := automateProgress(ev); msg != "" {
				emitter.emit(msg)
			}
			switch v := ev.AsAny().(type) {
			case tabstack.AutomateEventComplete:
				finalAnswer = v.Data.FinalAnswer
			case tabstack.AutomateEventError:
				return toolError(fmt.Errorf("automate: %s", v.Data.Error.Message)), nil, nil
			case tabstack.AutomateEventInteractiveFormDataRequest:
				// We can't prompt the user mid-tool-call. Auto-decline so the
				// agent doesn't hang for two minutes waiting for input.
				_, _ = c.Agent.AutomateInput(ctx, v.Data.RequestID, tabstack.AgentAutomateInputParams{
					Cancelled: param.NewOpt(true),
				})
				emitter.emit("Form-data request auto-declined (use the CLI for interactive mode)")
				_, _ = fmt.Fprintln(os.Stderr, "warning: tabstack_automate received an interactive:form_data:request event and auto-declined it (no interactive support in MCP yet)")
			}
		}
		if err := stream.Err(); err != nil {
			return toolError(fmt.Errorf("automate stream: %w", err)), nil, nil
		}
		if finalAnswer == "" {
			return toolError(fmt.Errorf("automate stream ended without a complete event")), nil, nil
		}
		return &sdk.CallToolResult{
			Content: []sdk.Content{&sdk.TextContent{Text: finalAnswer}},
		}, nil, nil
	})
}
