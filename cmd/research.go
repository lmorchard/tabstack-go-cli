package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/lmorchard/tabstack-go-cli/internal/client"
	"github.com/lmorchard/tabstack-go-cli/internal/spinner"
	"github.com/lmorchard/tabstack-go-cli/internal/sse"
	"github.com/spf13/cobra"
	tabstack "github.com/stainless-sdks/tabstack-go"
	"github.com/stainless-sdks/tabstack-go/packages/param"
)

var (
	researchMode         string
	researchNocache      bool
	researchFetchTimeout int64
	researchOutput       string
	researchColor        string
)

var researchCmd = &cobra.Command{
	Use:   "research <query>",
	Short: "Stream a multi-source AI research run",
	Args:  cobra.MaximumNArgs(1),
	RunE:  runResearch,
}

func runResearch(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return cmd.Help()
	}
	if err := validateEnum("mode", researchMode, validResearchModes); err != nil {
		return err
	}
	if err := validateEnum("output", researchOutput, validStreamOutputs); err != nil {
		return err
	}
	if err := validateEnum("color", researchColor, validColorModes); err != nil {
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
	if resolveStreamOutput(researchOutput) == "pretty" {
		defer func() { _ = stream.Close() }()
		started := time.Now()
		color := resolveStreamColor(researchColor)
		sp := spinner.New(os.Stdout, isTerminal(os.Stdout))
		sp.Start()
		defer sp.Stop()
		for stream.Next() {
			sp.ClearLine()
			if err := sse.PrettyResearch(os.Stdout, stream.Current(), started, color); err != nil {
				return fmt.Errorf("render event: %w", err)
			}
		}
		if err := stream.Err(); err != nil {
			return fmt.Errorf("research stream: %w", err)
		}
		return nil
	}
	if err := sse.WriteJSONLines[tabstack.ResearchEventUnion](os.Stdout, stream); err != nil {
		return fmt.Errorf("research stream: %w", err)
	}
	return nil
}

func init() {
	researchCmd.Flags().StringVar(&researchMode, "mode", "", "research mode: fast or balanced")
	researchCmd.Flags().BoolVar(&researchNocache, "nocache", false, "bypass cache")
	researchCmd.Flags().Int64Var(&researchFetchTimeout, "fetch-timeout", 0, "per-fetch timeout in seconds (0 = SDK default)")
	researchCmd.Flags().StringVar(&researchOutput, "output", "auto", "stream output format: auto (pretty on a TTY, json otherwise), json (one event per line), or pretty (human-readable)")
	researchCmd.Flags().StringVar(&researchColor, "color", "auto", "color in pretty output: auto (TTY only, respects NO_COLOR), always, or never")

	rootCmd.AddCommand(researchCmd)
}
