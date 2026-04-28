package mcp

import (
	"errors"
	"strings"
	"testing"

	sdk "github.com/modelcontextprotocol/go-sdk/mcp"
)

func TestJSONResult_PrettyPrintsValue(t *testing.T) {
	got := jsonResult(map[string]any{
		"title":   "Example Domain",
		"author":  "IANA",
		"visited": true,
	})
	if got.IsError {
		t.Fatal("IsError should be false on success")
	}
	if len(got.Content) != 1 {
		t.Fatalf("got %d content blocks, want 1", len(got.Content))
	}
	tc, ok := got.Content[0].(*sdk.TextContent)
	if !ok {
		t.Fatalf("expected TextContent, got %T", got.Content[0])
	}
	// MarshalIndent uses two-space indent and newlines; verify the shape.
	if !strings.Contains(tc.Text, `"title": "Example Domain"`) {
		t.Errorf("text missing title field: %q", tc.Text)
	}
	if !strings.Contains(tc.Text, "\n  ") {
		t.Errorf("expected indented JSON, got: %q", tc.Text)
	}
}

func TestToolError_SetsIsErrorAndUsesMessage(t *testing.T) {
	err := errors.New("API said no")
	got := toolError(err)
	if !got.IsError {
		t.Fatal("IsError should be true")
	}
	if len(got.Content) != 1 {
		t.Fatalf("got %d content blocks, want 1", len(got.Content))
	}
	tc, ok := got.Content[0].(*sdk.TextContent)
	if !ok {
		t.Fatalf("expected TextContent, got %T", got.Content[0])
	}
	if tc.Text != "API said no" {
		t.Errorf("text = %q, want %q", tc.Text, "API said no")
	}
}

// jsonResult on a value that fails to marshal (e.g. a chan) falls through to
// toolError instead of panicking.
func TestJSONResult_MarshalFailureBecomesToolError(t *testing.T) {
	got := jsonResult(map[string]any{"chan": make(chan int)})
	if !got.IsError {
		t.Fatal("expected IsError=true for non-marshalable input")
	}
}
