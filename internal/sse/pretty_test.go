package sse

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
	"time"

	tabstack "github.com/stainless-sdks/tabstack-go"
)

// fixedStart returns a time.Time the test can pass to elapsed() so the
// rendered duration prefix is deterministic.
func fixedStart() time.Time {
	return time.Now().Add(-12 * time.Second)
}

// automateFromJSON unmarshals a raw JSON event into AutomateEventUnion. The
// SDK's typed-variant accessors (AsAny() and friends) read from the union's
// internal raw JSON, so events have to be constructed via JSON unmarshal —
// setting the merged Data fields directly bypasses the type discriminator.
func automateFromJSON(t *testing.T, raw string) tabstack.AutomateEventUnion {
	t.Helper()
	var ev tabstack.AutomateEventUnion
	if err := json.Unmarshal([]byte(raw), &ev); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	return ev
}

func researchFromJSON(t *testing.T, raw string) tabstack.ResearchEventUnion {
	t.Helper()
	var ev tabstack.ResearchEventUnion
	if err := json.Unmarshal([]byte(raw), &ev); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	return ev
}

func TestPrettyAutomate_BrowserNavigated(t *testing.T) {
	ev := automateFromJSON(t, `{
		"event": "browser:navigated",
		"data": {
			"url": "https://example.com",
			"title": "Example Domain",
			"iterationId": "iter-1",
			"timestamp": 0
		}
	}`)
	var buf bytes.Buffer
	if err := PrettyAutomate(&buf, ev, fixedStart(), false); err != nil {
		t.Fatalf("PrettyAutomate: %v", err)
	}
	got := buf.String()
	if !strings.Contains(got, "browser:navigated https://example.com") {
		t.Errorf("output missing URL: %q", got)
	}
	if !strings.Contains(got, `"Example Domain"`) {
		t.Errorf("output missing title: %q", got)
	}
}

// agent:reasoned puts the full reasoning on a continuation line, preserving
// the original text verbatim — even when it's long enough that the previous
// (truncating) implementation would have cut it off.
func TestPrettyAutomate_AgentReasoned_FullText(t *testing.T) {
	long := "I need to look at the document title element. The title tag is in the HTML head. " +
		"Once I have the title, I can return it as the final answer. " +
		"The user is asking specifically for the page title, so I should be precise about that."
	ev := automateFromJSON(t, `{
		"event": "agent:reasoned",
		"data": {
			"reasoning": `+jsonString(long)+`,
			"iterationId": "iter-1",
			"timestamp": 0
		}
	}`)
	var buf bytes.Buffer
	if err := PrettyAutomate(&buf, ev, fixedStart(), false); err != nil {
		t.Fatalf("PrettyAutomate: %v", err)
	}
	got := buf.String()
	if !strings.Contains(got, long) {
		t.Errorf("reasoning text was modified or truncated:\nfull text not found in: %q", got)
	}
	if !strings.Contains(got, "agent:reasoned\n") {
		t.Errorf("reasoning should be on a continuation line: %q", got)
	}
}

func TestPrettyAutomate_AgentAction_RefAndValue(t *testing.T) {
	ev := automateFromJSON(t, `{
		"event": "agent:action",
		"data": {
			"action": "type",
			"ref": "E42",
			"value": "alice@example.com",
			"iterationId": "iter-1",
			"timestamp": 0
		}
	}`)
	var buf bytes.Buffer
	if err := PrettyAutomate(&buf, ev, fixedStart(), false); err != nil {
		t.Fatalf("PrettyAutomate: %v", err)
	}
	got := buf.String()
	if !strings.Contains(got, "ref=E42") {
		t.Errorf("missing ref: %q", got)
	}
	if !strings.Contains(got, `value="alice@example.com"`) {
		t.Errorf("missing value: %q", got)
	}
}

// "done"-style action with a multi-line value renders on a continuation line.
func TestPrettyAutomate_AgentAction_DoneMultiLine(t *testing.T) {
	val := "The page title is \"Example Domain\".\n\nThe full URL was https://example.com."
	ev := automateFromJSON(t, `{
		"event": "agent:action",
		"data": {
			"action": "done",
			"value": `+jsonString(val)+`,
			"iterationId": "iter-1",
			"timestamp": 0
		}
	}`)
	var buf bytes.Buffer
	if err := PrettyAutomate(&buf, ev, fixedStart(), false); err != nil {
		t.Fatalf("PrettyAutomate: %v", err)
	}
	got := buf.String()
	if !strings.Contains(got, "agent:action done\n") {
		t.Errorf("done action should put value on continuation line: %q", got)
	}
	if !strings.Contains(got, `The full URL was https://example.com.`) {
		t.Errorf("multi-line value not preserved: %q", got)
	}
}

func TestPrettyAutomate_Complete_SuccessFullAnswer(t *testing.T) {
	answer := "The page title is \"Example Domain\".\nIt's a placeholder domain reserved by IANA."
	ev := automateFromJSON(t, `{
		"event": "complete",
		"data": {
			"finalAnswer": `+jsonString(answer)+`,
			"success": true,
			"stats": {}
		}
	}`)
	var buf bytes.Buffer
	if err := PrettyAutomate(&buf, ev, fixedStart(), false); err != nil {
		t.Fatalf("PrettyAutomate: %v", err)
	}
	got := buf.String()
	if !strings.Contains(got, "complete ✓\n") {
		t.Errorf("expected success marker on header line: %q", got)
	}
	if !strings.Contains(got, "It's a placeholder domain reserved by IANA.") {
		t.Errorf("full answer not preserved: %q", got)
	}
}

func TestPrettyAutomate_Complete_Failure(t *testing.T) {
	ev := automateFromJSON(t, `{
		"event": "complete",
		"data": {
			"finalAnswer": "could not complete",
			"success": false,
			"stats": {}
		}
	}`)
	var buf bytes.Buffer
	if err := PrettyAutomate(&buf, ev, fixedStart(), false); err != nil {
		t.Fatalf("PrettyAutomate: %v", err)
	}
	got := buf.String()
	if !strings.Contains(got, "complete ✗") {
		t.Errorf("expected failure marker: %q", got)
	}
}

func TestPrettyAutomate_GenericFallback(t *testing.T) {
	ev := automateFromJSON(t, `{"event": "system:debug_message", "data": {"message": "x"}}`)
	var buf bytes.Buffer
	if err := PrettyAutomate(&buf, ev, fixedStart(), false); err != nil {
		t.Fatalf("PrettyAutomate: %v", err)
	}
	got := buf.String()
	if !strings.Contains(got, "system:debug_message") {
		t.Errorf("fallback should print event name: %q", got)
	}
}

func TestPrettyResearch_SearchingEnd(t *testing.T) {
	ev := researchFromJSON(t, `{
		"event": "searching:end",
		"data": {
			"iteration": 1,
			"message": "done searching",
			"urlsFound": 9,
			"urlsNew": 3,
			"timestamp": 0
		}
	}`)
	var buf bytes.Buffer
	if err := PrettyResearch(&buf, ev, fixedStart(), false); err != nil {
		t.Fatalf("PrettyResearch: %v", err)
	}
	got := buf.String()
	if !strings.Contains(got, "found 9 URL(s)") {
		t.Errorf("missing URL count: %q", got)
	}
	if !strings.Contains(got, "3 new") {
		t.Errorf("missing new count: %q", got)
	}
}

func TestPrettyResearch_PlanningEnd(t *testing.T) {
	ev := researchFromJSON(t, `{
		"event": "planning:end",
		"data": {
			"complexity": "simple",
			"message": "Plan ready",
			"objective": "x",
			"plan": "do x",
			"queries": ["q1", "q2", "q3"],
			"questions": [],
			"timestamp": 0
		}
	}`)
	var buf bytes.Buffer
	if err := PrettyResearch(&buf, ev, fixedStart(), false); err != nil {
		t.Fatalf("PrettyResearch: %v", err)
	}
	got := buf.String()
	if !strings.Contains(got, "complexity: simple") {
		t.Errorf("missing complexity: %q", got)
	}
	if !strings.Contains(got, "3 queries") {
		t.Errorf("missing query count: %q", got)
	}
}

// Research's complete event must preserve the full report verbatim — that's
// the actual research output, not progress noise. Multi-paragraph reports
// should pass through with paragraph breaks intact.
func TestPrettyResearch_Complete_PreservesFullReport(t *testing.T) {
	report := "Paris is the capital of France.\n\n" +
		"The city has been the political and cultural center of France for over a thousand years. " +
		"With a population of over 2 million people in the city proper, it serves as the seat of the French government.\n\n" +
		"Cited sources include encyclopedic references and government statistics."
	ev := researchFromJSON(t, `{
		"event": "complete",
		"data": {
			"message": "done",
			"metadata": {},
			"report": `+jsonString(report)+`,
			"timestamp": 0
		}
	}`)
	var buf bytes.Buffer
	if err := PrettyResearch(&buf, ev, fixedStart(), false); err != nil {
		t.Fatalf("PrettyResearch: %v", err)
	}
	got := buf.String()
	if !strings.Contains(got, "complete\n") {
		t.Errorf("complete header should be on its own line: %q", got)
	}
	// The full report (every paragraph) must be present.
	for _, frag := range []string{
		"Paris is the capital of France.",
		"With a population of over 2 million",
		"Cited sources include encyclopedic references",
	} {
		if !strings.Contains(got, frag) {
			t.Errorf("report fragment not preserved: %q\nfull output: %s", frag, got)
		}
	}
	// Paragraph breaks (\n\n) should survive.
	if !strings.Contains(got, "\n\n") {
		t.Errorf("paragraph breaks not preserved in report")
	}
}

func TestPrettyResearch_GenericFallback(t *testing.T) {
	ev := researchFromJSON(t, `{"event": "judging:start", "data": {"message": "x", "timestamp": 0}}`)
	var buf bytes.Buffer
	if err := PrettyResearch(&buf, ev, fixedStart(), false); err != nil {
		t.Fatalf("PrettyResearch: %v", err)
	}
	got := buf.String()
	if !strings.Contains(got, "judging:start") {
		t.Errorf("fallback should print event name: %q", got)
	}
}

// When color=true, output should contain ANSI SGR escape sequences. When
// color=false, the same event should produce plain text with no escapes.
// We don't assert the exact codes (that would over-fit the renderer's
// styling choices) — just the presence/absence of any escape.
func TestPrettyAutomate_ColorToggle(t *testing.T) {
	ev := automateFromJSON(t, `{
		"event": "complete",
		"data": {"finalAnswer": "all done", "success": true, "stats": {}}
	}`)

	var withColor, plain bytes.Buffer
	if err := PrettyAutomate(&withColor, ev, fixedStart(), true); err != nil {
		t.Fatalf("PrettyAutomate (color=true): %v", err)
	}
	if err := PrettyAutomate(&plain, ev, fixedStart(), false); err != nil {
		t.Fatalf("PrettyAutomate (color=false): %v", err)
	}

	const esc = "\x1b["
	if !strings.Contains(withColor.String(), esc) {
		t.Errorf("expected ANSI escape in colored output, got: %q", withColor.String())
	}
	if strings.Contains(plain.String(), esc) {
		t.Errorf("expected NO ANSI escape in plain output, got: %q", plain.String())
	}
	// The actual content (the success marker text and "all done") should be
	// present regardless of color.
	for _, frag := range []string{"complete", "all done"} {
		if !strings.Contains(plain.String(), frag) {
			t.Errorf("plain output missing %q: %q", frag, plain.String())
		}
		if !strings.Contains(withColor.String(), frag) {
			t.Errorf("colored output missing %q: %q", frag, withColor.String())
		}
	}
}

// oneLine collapses whitespace but never truncates.
func TestOneLine(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"short", "short"},
		{"  multi\n  line\n  text  ", "multi line text"},
		{"long\n\nstring\twith\tmixed\twhitespace", "long string with mixed whitespace"},
		{"", ""},
		{"   ", ""},
	}
	for _, tc := range tests {
		got := oneLine(tc.in)
		if got != tc.want {
			t.Errorf("oneLine(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

// jsonString quotes s as a JSON string literal — handy when constructing
// JSON test fixtures inline so embedded quotes/newlines don't break parsing.
func jsonString(s string) string {
	b, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	return string(b)
}
