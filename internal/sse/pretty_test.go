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
	if err := PrettyAutomate(&buf, ev, fixedStart()); err != nil {
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

func TestPrettyAutomate_AgentAction_PlainValue(t *testing.T) {
	ev := automateFromJSON(t, `{
		"event": "agent:action",
		"data": {
			"action": "done",
			"value": "The page title is Example Domain.",
			"iterationId": "iter-1",
			"timestamp": 0
		}
	}`)
	var buf bytes.Buffer
	_ = PrettyAutomate(&buf, ev, fixedStart())
	got := buf.String()
	if !strings.Contains(got, "agent:action done") {
		t.Errorf("missing action verb: %q", got)
	}
	if !strings.Contains(got, "The page title is Example Domain.") {
		t.Errorf("missing value: %q", got)
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
	_ = PrettyAutomate(&buf, ev, fixedStart())
	got := buf.String()
	if !strings.Contains(got, "ref=E42") {
		t.Errorf("missing ref: %q", got)
	}
	if !strings.Contains(got, `value="alice@example.com"`) {
		t.Errorf("missing value: %q", got)
	}
}

func TestPrettyAutomate_Complete_Success(t *testing.T) {
	ev := automateFromJSON(t, `{
		"event": "complete",
		"data": {
			"finalAnswer": "all done",
			"success": true,
			"stats": {}
		}
	}`)
	var buf bytes.Buffer
	_ = PrettyAutomate(&buf, ev, fixedStart())
	got := buf.String()
	if !strings.Contains(got, "complete ✓") {
		t.Errorf("expected success marker: %q", got)
	}
	if !strings.Contains(got, "all done") {
		t.Errorf("missing finalAnswer: %q", got)
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
	_ = PrettyAutomate(&buf, ev, fixedStart())
	got := buf.String()
	if !strings.Contains(got, "complete ✗") {
		t.Errorf("expected failure marker: %q", got)
	}
}

func TestPrettyAutomate_GenericFallback(t *testing.T) {
	ev := automateFromJSON(t, `{"event": "system:debug_message", "data": {"message": "x"}}`)
	var buf bytes.Buffer
	_ = PrettyAutomate(&buf, ev, fixedStart())
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
	_ = PrettyResearch(&buf, ev, fixedStart())
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
	_ = PrettyResearch(&buf, ev, fixedStart())
	got := buf.String()
	if !strings.Contains(got, "complexity: simple") {
		t.Errorf("missing complexity: %q", got)
	}
	if !strings.Contains(got, "3 queries") {
		t.Errorf("missing query count: %q", got)
	}
}

func TestPrettyResearch_GenericFallback(t *testing.T) {
	ev := researchFromJSON(t, `{"event": "judging:start", "data": {"message": "x", "timestamp": 0}}`)
	var buf bytes.Buffer
	_ = PrettyResearch(&buf, ev, fixedStart())
	got := buf.String()
	if !strings.Contains(got, "judging:start") {
		t.Errorf("fallback should print event name: %q", got)
	}
}

func TestTrunc(t *testing.T) {
	tests := []struct {
		in   string
		n    int
		want string
	}{
		{"short", 10, "short"},
		{"this is a long string that exceeds the limit", 20, "this is a long stri…"},
		{"  multi\n  line\n  text  ", 100, "multi line text"},
	}
	for _, tc := range tests {
		got := trunc(tc.in, tc.n)
		if got != tc.want {
			t.Errorf("trunc(%q, %d) = %q, want %q", tc.in, tc.n, got, tc.want)
		}
	}
}
