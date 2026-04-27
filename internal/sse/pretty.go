package sse

import (
	"fmt"
	"io"
	"strings"
	"time"

	tabstack "github.com/stainless-sdks/tabstack-go"
)

// truncWidth caps long strings (reasoning, reports, messages) when emitted
// inline. Full content is always available in the `--output json` mode.
const truncWidth = 140

// elapsed renders a coarse "[12s]" / "[1m23s]" prefix for the line.
func elapsed(start time.Time) string {
	return fmt.Sprintf("[%s]", time.Since(start).Round(time.Second))
}

// trunc returns s shortened to at most n runes, with an ellipsis appended if
// truncation occurred. Whitespace is collapsed so multi-line text fits on
// one terminal line.
func trunc(s string, n int) string {
	s = strings.TrimSpace(strings.Join(strings.Fields(s), " "))
	if len(s) <= n {
		return s
	}
	return s[:n-1] + "…"
}

// PrettyAutomate writes one human-readable line describing ev to w.
//
// The renderer recognizes the most informative event variants (browser
// navigation, agent step/reasoning/action, task lifecycle, complete, error)
// and renders them with their meaningful payload fields. Other variants fall
// through to a generic `[12s] event:name` line so nothing is silently
// dropped.
func PrettyAutomate(w io.Writer, ev tabstack.AutomateEventUnion, startedAt time.Time) error {
	prefix := elapsed(startedAt)
	switch v := ev.AsAny().(type) {
	case tabstack.AutomateEventCdpEndpointConnected:
		_, err := fmt.Fprintf(w, "%s cdp:endpoint_connected\n", prefix)
		return err
	case tabstack.AutomateEventAgentStatus:
		_, err := fmt.Fprintf(w, "%s agent:status — %s\n", prefix, trunc(v.Data.Message, truncWidth))
		return err
	case tabstack.AutomateEventAgentStep:
		_, err := fmt.Fprintf(w, "%s agent:step iteration %d\n", prefix, int(v.Data.CurrentIteration))
		return err
	case tabstack.AutomateEventAgentReasoned:
		_, err := fmt.Fprintf(w, "%s agent:reasoned — %s\n", prefix, trunc(v.Data.Reasoning, truncWidth))
		return err
	case tabstack.AutomateEventAgentAction:
		switch {
		case v.Data.Value != "" && v.Data.Ref != "":
			_, err := fmt.Fprintf(w, "%s agent:action %s ref=%s value=%q\n", prefix, v.Data.Action, v.Data.Ref, trunc(v.Data.Value, truncWidth))
			return err
		case v.Data.Value != "":
			_, err := fmt.Fprintf(w, "%s agent:action %s — %s\n", prefix, v.Data.Action, trunc(v.Data.Value, truncWidth))
			return err
		default:
			_, err := fmt.Fprintf(w, "%s agent:action %s\n", prefix, v.Data.Action)
			return err
		}
	case tabstack.AutomateEventBrowserNavigated:
		_, err := fmt.Fprintf(w, "%s browser:navigated %s — %q\n", prefix, v.Data.URL, trunc(v.Data.Title, truncWidth))
		return err
	case tabstack.AutomateEventBrowserActionStarted:
		_, err := fmt.Fprintf(w, "%s browser:action_started\n", prefix)
		return err
	case tabstack.AutomateEventBrowserActionCompleted:
		_, err := fmt.Fprintf(w, "%s browser:action_completed\n", prefix)
		return err
	case tabstack.AutomateEventTaskStarted:
		_, err := fmt.Fprintf(w, "%s task:started\n", prefix)
		return err
	case tabstack.AutomateEventTaskValidated:
		_, err := fmt.Fprintf(w, "%s task:validated\n", prefix)
		return err
	case tabstack.AutomateEventTaskCompleted:
		_, err := fmt.Fprintf(w, "%s task:completed\n", prefix)
		return err
	case tabstack.AutomateEventTaskAborted:
		_, err := fmt.Fprintf(w, "%s task:aborted\n", prefix)
		return err
	case tabstack.AutomateEventInteractiveFormDataRequest:
		_, err := fmt.Fprintf(w, "%s interactive:form_data:request — %s (%d field(s))\n",
			prefix, trunc(v.Data.FormDescription, truncWidth-30), len(v.Data.Fields))
		return err
	case tabstack.AutomateEventComplete:
		mark := "✓"
		if !v.Data.Success {
			mark = "✗"
		}
		_, err := fmt.Fprintf(w, "%s complete %s %q\n", prefix, mark, trunc(v.Data.FinalAnswer, truncWidth))
		return err
	case tabstack.AutomateEventError:
		_, err := fmt.Fprintf(w, "%s error — %s\n", prefix, trunc(v.Data.Error.Message, truncWidth))
		return err
	default:
		// Unknown / unspecialized variant. Fall back to event-name only.
		_, err := fmt.Fprintf(w, "%s %s\n", prefix, ev.Event)
		return err
	}
}

// PrettyResearch writes one human-readable line describing ev to w.
//
// Most research event variants carry a `Message` field with a human-friendly
// description; for those we render `[12s] event:name — <message>`. A few
// variants get extra fields (search counts, complexity, etc.).
func PrettyResearch(w io.Writer, ev tabstack.ResearchEventUnion, startedAt time.Time) error {
	prefix := elapsed(startedAt)
	switch v := ev.AsAny().(type) {
	case tabstack.ResearchEventPlanningStart:
		_, err := fmt.Fprintf(w, "%s planning:start — %s\n", prefix, trunc(v.Data.Message, truncWidth))
		return err
	case tabstack.ResearchEventPlanningEnd:
		_, err := fmt.Fprintf(w, "%s planning:end — %s (complexity: %s, %d queries)\n",
			prefix, trunc(v.Data.Message, 60), v.Data.Complexity, len(v.Data.Queries))
		return err
	case tabstack.ResearchEventIterationStart:
		_, err := fmt.Fprintf(w, "%s iteration:start [%d/%d] — %s\n",
			prefix, int(v.Data.Iteration), int(v.Data.MaxIterations), trunc(v.Data.Message, 80))
		return err
	case tabstack.ResearchEventIterationEnd:
		stop := v.Data.StopReason
		if stop == "" {
			stop = "—"
		}
		_, err := fmt.Fprintf(w, "%s iteration:end [%d] (stop: %s)\n", prefix, int(v.Data.Iteration), stop)
		return err
	case tabstack.ResearchEventSearchingStart:
		_, err := fmt.Fprintf(w, "%s searching:start — %d quer%s\n",
			prefix, len(v.Data.Queries), pluralY(len(v.Data.Queries)))
		return err
	case tabstack.ResearchEventSearchingEnd:
		_, err := fmt.Fprintf(w, "%s searching:end — found %d URL(s), %d new\n",
			prefix, int(v.Data.URLsFound), int(v.Data.URLsNew))
		return err
	case tabstack.ResearchEventWritingStart:
		_, err := fmt.Fprintf(w, "%s writing:start — attempt %d/%d\n",
			prefix, int(v.Data.Attempt), int(v.Data.MaxAttempts))
		return err
	case tabstack.ResearchEventWritingEnd:
		_, err := fmt.Fprintf(w, "%s writing:end — attempt %d\n", prefix, int(v.Data.Attempt))
		return err
	case tabstack.ResearchEventComplete:
		_, err := fmt.Fprintf(w, "%s complete — %s\n", prefix, trunc(v.Data.Report, truncWidth))
		return err
	case tabstack.ResearchEventError:
		_, err := fmt.Fprintf(w, "%s error — %s\n", prefix, trunc(v.Data.Error.Message, truncWidth))
		return err
	default:
		// Unknown / unspecialized variant. Fall back to event-name only.
		_, err := fmt.Fprintf(w, "%s %s\n", prefix, ev.Event)
		return err
	}
}

func pluralY(n int) string {
	if n == 1 {
		return "y"
	}
	return "ies"
}
