package mcp

import (
	"encoding/json"

	sdk "github.com/modelcontextprotocol/go-sdk/mcp"
)

// jsonResult serializes v as pretty-printed JSON and wraps it in a single
// TextContent block. Falls back to a tool error if marshaling fails (which
// shouldn't happen for SDK response types but isn't worth panicking over).
func jsonResult(v any) *sdk.CallToolResult {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return toolError(err)
	}
	return &sdk.CallToolResult{
		Content: []sdk.Content{&sdk.TextContent{Text: string(b)}},
	}
}

// textResult wraps s in a single TextContent block — used for streaming-tool
// results where the payload is already a plain string (the report or final
// answer).
func textResult(s string) *sdk.CallToolResult {
	return &sdk.CallToolResult{
		Content: []sdk.Content{&sdk.TextContent{Text: s}},
	}
}

// toolError wraps a runtime error as a non-protocol tool error: IsError=true,
// message in TextContent. The MCP SDK distinguishes this from a returned
// `error` value, which becomes a JSON-RPC protocol-level error (wrong for
// "tool ran but the API said no").
func toolError(err error) *sdk.CallToolResult {
	return &sdk.CallToolResult{
		Content: []sdk.Content{&sdk.TextContent{Text: err.Error()}},
		IsError: true,
	}
}
