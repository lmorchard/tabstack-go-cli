package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/lmorchard/tabstack-go-cli/internal/client"
	"github.com/lmorchard/tabstack-go-cli/internal/sse"
	"github.com/spf13/cobra"
	tabstack "github.com/stainless-sdks/tabstack-go"
	"github.com/stainless-sdks/tabstack-go/packages/param"
)

var (
	researchMode         string
	researchNocache      bool
	researchFetchTimeout int64
)

var researchCmd = &cobra.Command{
	Use:   "research <query>",
	Short: "Stream a multi-source AI research run",
	Args:  cobra.ExactArgs(1),
	RunE:  runResearch,
}

func runResearch(_ *cobra.Command, args []string) error {
	if err := validateEnum("mode", researchMode, validResearchModes); err != nil {
		return err
	}
	c, err := client.New(GetConfig())
	if err != nil {
		return err
	}

	body := tabstack.AgentResearchParams{
		Query: args[0],
	}
	if researchMode != "" {
		body.Mode = tabstack.AgentResearchParamsMode(researchMode)
	}
	if researchNocache {
		body.Nocache = param.NewOpt(true)
	}
	if researchFetchTimeout > 0 {
		body.FetchTimeout = param.NewOpt(researchFetchTimeout)
	}

	stream := c.Agent.ResearchStreaming(context.Background(), body)
	if err := sse.WriteJSONLines[tabstack.ResearchEvent](os.Stdout, stream); err != nil {
		return fmt.Errorf("research stream: %w", err)
	}
	return nil
}

func init() {
	researchCmd.Flags().StringVar(&researchMode, "mode", "", "research mode: fast or balanced")
	researchCmd.Flags().BoolVar(&researchNocache, "nocache", false, "bypass cache")
	researchCmd.Flags().Int64Var(&researchFetchTimeout, "fetch-timeout", 0, "per-fetch timeout in seconds (0 = SDK default)")

	rootCmd.AddCommand(researchCmd)
}
