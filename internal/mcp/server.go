// Package mcp embeds an MCP (Model Context Protocol) server inside the
// tabstack CLI binary. The server exposes every Tabstack API operation as
// an MCP tool, reusing the existing internal/client and the vendored
// Stainless-generated SDK.
package mcp

import (
	sdk "github.com/modelcontextprotocol/go-sdk/mcp"
	tabstack "github.com/stainless-sdks/tabstack-go"
)

// NewServer constructs an MCP server preloaded with all Tabstack tools.
//
// `version` is the binary's version string (e.g. set via ldflags). MCP
// clients use Implementation.Version to render the server identity.
func NewServer(client *tabstack.Client, version string) *sdk.Server {
	s := sdk.NewServer(&sdk.Implementation{
		Name:    "tabstack",
		Version: version,
	}, nil)
	// Tool registrations land here in Phase 3+.
	_ = client
	return s
}
