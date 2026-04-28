package mcp

import (
	"encoding/json"
	"strings"
	"testing"

	tabstack "github.com/stainless-sdks/tabstack-go"
)

// automateFromJSON unmarshals a raw JSON event into AutomateEventUnion. The
// SDK's AsAny() accessor reads from the union's internal raw JSON, so events
// have to be constructed via JSON unmarshal — populating the merged Data
// fields directly bypasses the type discriminator.
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

func TestAutomateProgress(t *testing.T) {
	cases := []struct {
		name     string
		raw      string
		want     string // empty means "should be skipped"
		contains string // alternative to want — just verify substring
	}{
		{
			name: "cdp:endpoint_connected",
			raw:  `{"event":"cdp:endpoint_connected","data":{}}`,
			want: "Connected to browser",
		},
		{
			name: "browser:navigated includes URL",
			raw:  `{"event":"browser:navigated","data":{"url":"https://example.com","title":"Ex","iterationId":"i1","timestamp":0}}`,
			want: "Navigated to https://example.com",
		},
		{
			name:     "agent:reasoned truncates and prefixes",
			raw:      `{"event":"agent:reasoned","data":{"reasoning":"some long thought","iterationId":"i1","timestamp":0}}`,
			contains: "Reasoning: some long thought",
		},
		{
			name: "task:completed",
			raw:  `{"event":"task:completed","data":{}}`,
			want: "Task completed",
		},
		{
			name: "complete",
			raw:  `{"event":"complete","data":{"finalAnswer":"x","success":true,"stats":{}}}`,
			want: "Done",
		},
		{
			name: "noise event is skipped",
			raw:  `{"event":"system:debug_message","data":{"message":"x"}}`,
			want: "",
		},
		{
			name: "another noise event is skipped",
			raw:  `{"event":"cdp:endpoint_cycle","data":{}}`,
			want: "",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ev := automateFromJSON(t, tc.raw)
			got := automateProgress(ev)
			switch {
			case tc.want != "" && got != tc.want:
				t.Errorf("got %q, want %q", got, tc.want)
			case tc.contains != "" && !strings.Contains(got, tc.contains):
				t.Errorf("got %q, want it to contain %q", got, tc.contains)
			case tc.want == "" && tc.contains == "" && got != "":
				t.Errorf("expected empty (skipped), got %q", got)
			}
		})
	}
}

// agent:reasoned with an absurdly long reasoning string gets capped at ~120 runes
// plus the ellipsis. Rune-safe — multi-byte chars don't get cut mid-sequence.
func TestAutomateProgress_AgentReasonedTruncates(t *testing.T) {
	long := strings.Repeat("a", 200)
	ev := automateFromJSON(t, `{"event":"agent:reasoned","data":{"reasoning":"`+long+`","iterationId":"i","timestamp":0}}`)
	got := automateProgress(ev)
	if !strings.HasPrefix(got, "Reasoning: ") {
		t.Errorf("missing prefix: %q", got)
	}
	if !strings.HasSuffix(got, "…") {
		t.Errorf("expected truncation ellipsis: %q", got)
	}
	// Length ≈ 11 ("Reasoning: ") + 119 + 1 (ellipsis) = 131 runes.
	if runes := []rune(got); len(runes) > 135 {
		t.Errorf("output too long (%d runes): %q", len(runes), got)
	}
}

func TestResearchProgress(t *testing.T) {
	cases := []struct {
		name     string
		raw      string
		want     string
		contains string
	}{
		{
			name:     "planning:end with complexity and query count",
			raw:      `{"event":"planning:end","data":{"complexity":"simple","message":"x","objective":"x","plan":"x","queries":["q1","q2"],"questions":[],"timestamp":0}}`,
			contains: "Plan ready (simple, 2 queries)",
		},
		{
			name:     "searching:end counts URLs",
			raw:      `{"event":"searching:end","data":{"iteration":1,"message":"x","urlsFound":9,"urlsNew":3,"timestamp":0}}`,
			contains: "Found 9 URL(s), 3 new",
		},
		{
			name:     "searching:start with single query uses singular",
			raw:      `{"event":"searching:start","data":{"iteration":1,"message":"x","queries":["q"],"timestamp":0}}`,
			contains: "1 query",
		},
		{
			name:     "searching:start with multiple uses plural",
			raw:      `{"event":"searching:start","data":{"iteration":1,"message":"x","queries":["q1","q2","q3"],"timestamp":0}}`,
			contains: "3 queries",
		},
		{
			name: "noise event is skipped",
			raw:  `{"event":"judging:start","data":{"message":"x","timestamp":0}}`,
			want: "",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ev := researchFromJSON(t, tc.raw)
			got := researchProgress(ev)
			switch {
			case tc.want != "" && got != tc.want:
				t.Errorf("got %q, want %q", got, tc.want)
			case tc.contains != "" && !strings.Contains(got, tc.contains):
				t.Errorf("got %q, want it to contain %q", got, tc.contains)
			case tc.want == "" && tc.contains == "" && got != "":
				t.Errorf("expected empty (skipped), got %q", got)
			}
		})
	}
}
