package sse

import (
	"fmt"
	"io"
	"strings"
	"time"
	"unicode"

	tabstack "github.com/stainless-sdks/tabstack-go"
)

// elapsed renders a coarse "[12s]" / "[1m23s]" prefix for the line.
func elapsed(start time.Time) string {
	return fmt.Sprintf("[%s]", time.Since(start).Round(time.Second))
}

// sanitize strips control characters from API-supplied strings, keeping `\n`
// and `\t` so multi-line content is preserved. Without this an attacker-
// controlled response could inject ANSI escape sequences (rewriting earlier
// terminal lines, changing the window title, etc.) — even when the renderer
// runs with --color=never, since the dangerous sequences would come from the
// payload, not the styler.
func sanitize(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	for _, r := range s {
		switch {
		case r == '\n' || r == '\t':
			b.WriteRune(r)
		case unicode.IsControl(r):
			// drop
		default:
			b.WriteRune(r)
		}
	}
	return b.String()
}

// content prepares an API-supplied multi-line string for terminal display:
// sanitizes control characters (preventing escape injection) and trims a
// trailing newline so the renderer's own newline doesn't double up. Internal
// newlines, tabs, and leading whitespace are preserved verbatim.
func content(s string) string {
	return strings.TrimRight(sanitize(s), "\n")
}

// oneLine collapses internal whitespace (including newlines) into single
// spaces and strips control characters. Used for fields that render inline
// in the event header — multi-line content there would break the prefix
// layout. Does NOT truncate; full content is preserved, just on one line.
func oneLine(s string) string {
	return strings.TrimSpace(strings.Join(strings.Fields(sanitize(s)), " "))
}

// style applies ANSI SGR escapes when enabled. Hand-rolled rather than
// pulling in a color library — we use a tiny set of codes (dim, bold, three
// foreground colors) and the SGR sequences are stable.
type style struct{ enabled bool }

func (s style) wrap(code, text string) string {
	if !s.enabled {
		return text
	}
	return "\x1b[" + code + "m" + text + "\x1b[0m"
}

func (s style) dim(t string) string       { return s.wrap("2", t) }
func (s style) bold(t string) string      { return s.wrap("1", t) }
func (s style) green(t string) string     { return s.wrap("32", t) }
func (s style) red(t string) string       { return s.wrap("31", t) }
func (s style) cyan(t string) string      { return s.wrap("36", t) }
func (s style) underline(t string) string { return s.wrap("4", t) }

// header renders the standard "[Xs] event:name" header chunk with the prefix
// dimmed and the event name in cyan when color is enabled.
func (s style) header(prefix, event string) string {
	return s.dim(prefix) + " " + s.cyan(event)
}

// PrettyAutomate writes one or more human-readable lines describing ev to w.
//
// Most variants render as a single header line. Variants that carry
// substantial output — agent:reasoned, complete (final answer), error
// (message) — print a header line followed by the full content on a
// continuation line, preserving newlines so paragraph structure survives.
// Unspecialized variants fall through to a generic `[12s] event:name` line
// so nothing is silently dropped.
//
// When color is true, ANSI SGR escapes highlight the time prefix (dim),
// event names (cyan), success/failure markers (green/red), and error event
// names (red). Caller is responsible for deciding whether color is
// appropriate (TTY check, NO_COLOR env, etc.).
func PrettyAutomate(w io.Writer, ev tabstack.AutomateEventUnion, startedAt time.Time, color bool) error {
	prefix := elapsed(startedAt)
	s := style{enabled: color}

	switch v := ev.AsAny().(type) {
	case tabstack.AutomateEventCdpEndpointConnected:
		_, err := fmt.Fprintf(w, "%s\n", s.header(prefix, "cdp:endpoint_connected"))
		return err
	case tabstack.AutomateEventAgentStatus:
		_, err := fmt.Fprintf(w, "%s — %s\n", s.header(prefix, "agent:status"), oneLine(v.Data.Message))
		return err
	case tabstack.AutomateEventAgentStep:
		_, err := fmt.Fprintf(w, "%s iteration %d\n", s.header(prefix, "agent:step"), int(v.Data.CurrentIteration))
		return err
	case tabstack.AutomateEventAgentReasoned:
		_, err := fmt.Fprintf(w, "%s\n%s\n", s.header(prefix, "agent:reasoned"), content(v.Data.Reasoning))
		return err
	case tabstack.AutomateEventAgentAction:
		switch {
		case v.Data.Ref != "" && v.Data.Value != "":
			_, err := fmt.Fprintf(w, "%s %s ref=%s value=%q\n",
				s.header(prefix, "agent:action"), s.bold(v.Data.Action), v.Data.Ref, oneLine(v.Data.Value))
			return err
		case v.Data.Ref != "":
			_, err := fmt.Fprintf(w, "%s %s ref=%s\n",
				s.header(prefix, "agent:action"), s.bold(v.Data.Action), v.Data.Ref)
			return err
		case v.Data.Value != "":
			// "done"-style action whose value can be a multi-paragraph answer.
			_, err := fmt.Fprintf(w, "%s %s\n%s\n",
				s.header(prefix, "agent:action"), s.bold(v.Data.Action), content(v.Data.Value))
			return err
		default:
			_, err := fmt.Fprintf(w, "%s %s\n", s.header(prefix, "agent:action"), s.bold(v.Data.Action))
			return err
		}
	case tabstack.AutomateEventBrowserNavigated:
		_, err := fmt.Fprintf(w, "%s %s — %q\n",
			s.header(prefix, "browser:navigated"), s.underline(v.Data.URL), oneLine(v.Data.Title))
		return err
	case tabstack.AutomateEventBrowserActionStarted:
		_, err := fmt.Fprintf(w, "%s\n", s.header(prefix, "browser:action_started"))
		return err
	case tabstack.AutomateEventBrowserActionCompleted:
		_, err := fmt.Fprintf(w, "%s\n", s.header(prefix, "browser:action_completed"))
		return err
	case tabstack.AutomateEventTaskStarted:
		_, err := fmt.Fprintf(w, "%s\n", s.header(prefix, "task:started"))
		return err
	case tabstack.AutomateEventTaskValidated:
		_, err := fmt.Fprintf(w, "%s\n", s.header(prefix, "task:validated"))
		return err
	case tabstack.AutomateEventTaskCompleted:
		_, err := fmt.Fprintf(w, "%s\n", s.header(prefix, "task:completed"))
		return err
	case tabstack.AutomateEventTaskAborted:
		_, err := fmt.Fprintf(w, "%s\n", s.header(prefix, "task:aborted"))
		return err
	case tabstack.AutomateEventInteractiveFormDataRequest:
		_, err := fmt.Fprintf(w, "%s — %s (%d field(s))\n",
			s.header(prefix, "interactive:form_data:request"), oneLine(v.Data.FormDescription), len(v.Data.Fields))
		return err
	case tabstack.AutomateEventComplete:
		var mark string
		if v.Data.Success {
			mark = s.green("✓")
		} else {
			mark = s.red("✗")
		}
		if v.Data.FinalAnswer != "" {
			_, err := fmt.Fprintf(w, "%s %s\n%s\n",
				s.header(prefix, "complete"), mark, content(v.Data.FinalAnswer))
			return err
		}
		_, err := fmt.Fprintf(w, "%s %s\n", s.header(prefix, "complete"), mark)
		return err
	case tabstack.AutomateEventError:
		_, err := fmt.Fprintf(w, "%s %s\n%s\n",
			s.dim(prefix), s.red("error"), content(v.Data.Error.Message))
		return err
	default:
		// Unknown / unspecialized variant. Fall back to event-name only.
		_, err := fmt.Fprintf(w, "%s\n", s.header(prefix, ev.Event))
		return err
	}
}

// PrettyResearch writes one or more human-readable lines describing ev to w.
//
// Same conventions as PrettyAutomate: terminal-result variants (complete,
// error) print the full payload on a continuation line; progress variants
// render on a single line. Color usage matches PrettyAutomate.
func PrettyResearch(w io.Writer, ev tabstack.ResearchEventUnion, startedAt time.Time, color bool) error {
	prefix := elapsed(startedAt)
	s := style{enabled: color}

	switch v := ev.AsAny().(type) {
	case tabstack.ResearchEventPlanningStart:
		_, err := fmt.Fprintf(w, "%s — %s\n", s.header(prefix, "planning:start"), oneLine(v.Data.Message))
		return err
	case tabstack.ResearchEventPlanningEnd:
		_, err := fmt.Fprintf(w, "%s — %s (complexity: %s, %d queries)\n",
			s.header(prefix, "planning:end"), oneLine(v.Data.Message), v.Data.Complexity, len(v.Data.Queries))
		return err
	case tabstack.ResearchEventIterationStart:
		_, err := fmt.Fprintf(w, "%s [%d/%d] — %s\n",
			s.header(prefix, "iteration:start"), int(v.Data.Iteration), int(v.Data.MaxIterations), oneLine(v.Data.Message))
		return err
	case tabstack.ResearchEventIterationEnd:
		stop := v.Data.StopReason
		if stop == "" {
			stop = "—"
		}
		_, err := fmt.Fprintf(w, "%s [%d] (stop: %s)\n",
			s.header(prefix, "iteration:end"), int(v.Data.Iteration), stop)
		return err
	case tabstack.ResearchEventSearchingStart:
		_, err := fmt.Fprintf(w, "%s — %d quer%s\n",
			s.header(prefix, "searching:start"), len(v.Data.Queries), pluralY(len(v.Data.Queries)))
		return err
	case tabstack.ResearchEventSearchingEnd:
		_, err := fmt.Fprintf(w, "%s — found %d URL(s), %d new\n",
			s.header(prefix, "searching:end"), int(v.Data.URLsFound), int(v.Data.URLsNew))
		return err
	case tabstack.ResearchEventWritingStart:
		_, err := fmt.Fprintf(w, "%s — attempt %d/%d\n",
			s.header(prefix, "writing:start"), int(v.Data.Attempt), int(v.Data.MaxAttempts))
		return err
	case tabstack.ResearchEventWritingEnd:
		_, err := fmt.Fprintf(w, "%s — attempt %d\n", s.header(prefix, "writing:end"), int(v.Data.Attempt))
		return err
	case tabstack.ResearchEventComplete:
		_, err := fmt.Fprintf(w, "%s\n%s\n", s.header(prefix, "complete"), content(v.Data.Report))
		return err
	case tabstack.ResearchEventError:
		_, err := fmt.Fprintf(w, "%s %s\n%s\n",
			s.dim(prefix), s.red("error"), content(v.Data.Error.Message))
		return err
	default:
		// Unknown / unspecialized variant. Fall back to event-name only.
		_, err := fmt.Fprintf(w, "%s\n", s.header(prefix, ev.Event))
		return err
	}
}

func pluralY(n int) string {
	if n == 1 {
		return "y"
	}
	return "ies"
}
