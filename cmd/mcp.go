package cmd

import (
	"fmt"

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
	return fmt.Errorf("not yet implemented (Task 4 wires the server)")
}

func init() {
	mcpCmd.Flags().StringVar(&mcpTransport, "transport", "stdio",
		"wire protocol: stdio (for Claude Code etc.) or http")
	mcpCmd.Flags().StringVar(&mcpListen, "listen", "127.0.0.1:8080",
		"HTTP bind address (used only when --transport http). Localhost-only by default.")
	rootCmd.AddCommand(mcpCmd)
}
