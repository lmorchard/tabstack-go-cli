package cmd

import (
	"context"
	"fmt"
	"net/http"

	"github.com/lmorchard/tabstack-go-cli/internal/client"
	mcppkg "github.com/lmorchard/tabstack-go-cli/internal/mcp"
	sdk "github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/spf13/cobra"
)

var (
	mcpTransport string
	mcpListen    string
)

var mcpCmd = &cobra.Command{
	Use:   "mcp",
	Short: "Run the Tabstack API as an MCP (Model Context Protocol) server",
	Long: `Expose every Tabstack API operation as an MCP tool, usable from
Claude Code, Cursor, or any other MCP-aware client. The server reuses the
same TABSTACK_API_KEY env / --api-key flag / tabstack.yaml config as the
other CLI subcommands.`,
	RunE: runMCP,
}

func runMCP(_ *cobra.Command, _ []string) error {
	if err := validateEnum("transport", mcpTransport, validMCPTransports); err != nil {
		return err
	}

	c, err := client.New(GetConfig())
	if err != nil {
		return err
	}
	srv := mcppkg.NewServer(c, version)
	ctx := context.Background()

	switch mcpTransport {
	case "stdio":
		GetLogger().Info("MCP server starting on stdio")
		return srv.Run(ctx, &sdk.StdioTransport{})
	case "http":
		GetLogger().Infof("MCP server starting on http://%s", mcpListen)
		handler := sdk.NewStreamableHTTPHandler(func(*http.Request) *sdk.Server { return srv }, nil)
		return http.ListenAndServe(mcpListen, handler) //nolint:gosec // localhost-only by default; external bind is a documented opt-in
	}
	return fmt.Errorf("unreachable: validateEnum should have caught %q", mcpTransport)
}

func init() {
	mcpCmd.Flags().StringVar(&mcpTransport, "transport", "stdio",
		"wire protocol: stdio (for Claude Code etc.) or http")
	mcpCmd.Flags().StringVar(&mcpListen, "listen", "127.0.0.1:8080",
		"HTTP bind address (used only when --transport http). Localhost-only by default.")
	rootCmd.AddCommand(mcpCmd)
}
