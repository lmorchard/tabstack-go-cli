package mcp

import (
	"context"
	"fmt"
	"strings"

	sdk "github.com/modelcontextprotocol/go-sdk/mcp"
	tabstack "github.com/stainless-sdks/tabstack-go"
)

// progressEmitter sends MCP progress notifications during a streaming tool
// call. Created at the top of a streaming handler; emit() is called per
// curated event. If the client didn't supply a progress token, all calls
// are no-ops (progress notifications would be unreachable anyway).
type progressEmitter struct {
	ctx     context.Context
	session *sdk.ServerSession
	token   any
	count   float64
}

func newProgressEmitter(ctx context.Context, req *sdk.CallToolRequest) *progressEmitter {
	return &progressEmitter{
		ctx:     ctx,
		session: req.Session,
		token:   req.Params.GetProgressToken(),
	}
}

// emit sends a progress notification with the given message. Errors are
// swallowed — a notification failing shouldn't fail the tool call.
func (p *progressEmitter) emit(msg string) {
	if p.token == nil {
		return
	}
	p.count++
	_ = p.session.NotifyProgress(p.ctx, &sdk.ProgressNotificationParams{
		Message:       msg,
		ProgressToken: p.token,
		Progress:      p.count,
	})
}

// automateProgress returns the progress message for the given event variant,
// or "" if the event should be skipped (noise events: cdp:endpoint_cycle,
// system:debug_*, ai:generation chunks). Mirrors the curated subset that
// internal/sse's pretty renderer specializes.
func automateProgress(ev tabstack.AutomateEventUnion) string {
	switch v := ev.AsAny().(type) {
	case tabstack.AutomateEventCdpEndpointConnected:
		return "Connected to browser"
	case tabstack.AutomateEventAgentStatus:
		return v.Data.Message
	case tabstack.AutomateEventAgentStep:
		return fmt.Sprintf("Step %d", int(v.Data.CurrentIteration))
	case tabstack.AutomateEventBrowserNavigated:
		return fmt.Sprintf("Navigated to %s", v.Data.URL)
	case tabstack.AutomateEventAgentAction:
		return fmt.Sprintf("Action: %s", v.Data.Action)
	case tabstack.AutomateEventAgentReasoned:
		// Reasoning can be many paragraphs; the progress message is for a
		// status-line UI. Collapse whitespace and cap length so it fits.
		// Full reasoning stays in the SSE stream and would be available in
		// a future "all events" tool result.
		msg := strings.Join(strings.Fields(v.Data.Reasoning), " ")
		runes := []rune(msg)
		if len(runes) > 120 {
			msg = string(runes[:119]) + "…"
		}
		return "Reasoning: " + msg
	case tabstack.AutomateEventTaskStarted:
		return "Task started"
	case tabstack.AutomateEventTaskValidated:
		return "Task validated"
	case tabstack.AutomateEventTaskCompleted:
		return "Task completed"
	case tabstack.AutomateEventTaskAborted:
		return "Task aborted"
	case tabstack.AutomateEventComplete:
		return "Done"
	case tabstack.AutomateEventError:
		return "Error: " + v.Data.Error.Message
	default:
		return ""
	}
}

// researchProgress returns the progress message for the given event variant,
// or "" if the event should be skipped.
func researchProgress(ev tabstack.ResearchEventUnion) string {
	switch v := ev.AsAny().(type) {
	case tabstack.ResearchEventPlanningStart:
		return "Planning research"
	case tabstack.ResearchEventPlanningEnd:
		return fmt.Sprintf("Plan ready (%s, %d queries)", v.Data.Complexity, len(v.Data.Queries))
	case tabstack.ResearchEventIterationStart:
		return fmt.Sprintf("Iteration %d/%d", int(v.Data.Iteration), int(v.Data.MaxIterations))
	case tabstack.ResearchEventSearchingStart:
		return fmt.Sprintf("Searching (%d quer%s)", len(v.Data.Queries), pluralYY(len(v.Data.Queries)))
	case tabstack.ResearchEventSearchingEnd:
		return fmt.Sprintf("Found %d URL(s), %d new", int(v.Data.URLsFound), int(v.Data.URLsNew))
	case tabstack.ResearchEventWritingStart:
		return fmt.Sprintf("Writing report (attempt %d/%d)", int(v.Data.Attempt), int(v.Data.MaxAttempts))
	case tabstack.ResearchEventWritingEnd:
		return "Report draft complete"
	case tabstack.ResearchEventComplete:
		return "Done"
	case tabstack.ResearchEventError:
		return "Error: " + v.Data.Error.Message
	default:
		return ""
	}
}

func pluralYY(n int) string {
	if n == 1 {
		return "y"
	}
	return "ies"
}
