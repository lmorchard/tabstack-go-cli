package mcp

import (
	"context"
	"fmt"

	sdk "github.com/modelcontextprotocol/go-sdk/mcp"
	tabstack "github.com/stainless-sdks/tabstack-go"
	"github.com/stainless-sdks/tabstack-go/packages/param"
)

// ResearchInput is the input for tabstack_research.
type ResearchInput struct {
	Query        string `json:"query" jsonschema:"the research question or query to answer"`
	Mode         string `json:"mode,omitempty" jsonschema:"research mode: fast (single iteration, default) or balanced (multiple iterations)"`
	Nocache      bool   `json:"nocache,omitempty" jsonschema:"bypass cache and force fresh research"`
	FetchTimeout int64  `json:"fetch_timeout,omitempty" jsonschema:"per-fetch timeout in seconds (0 = SDK default)"`
}

func registerResearch(s *sdk.Server, c *tabstack.Client) {
	sdk.AddTool(s, &sdk.Tool{
		Name: "tabstack_research",
		Description: "Run a multi-source AI research query and return a cited markdown report. " +
			"Use for questions that require synthesizing information from multiple web sources.",
	}, func(ctx context.Context, req *sdk.CallToolRequest, in *ResearchInput) (*sdk.CallToolResult, any, error) {
		body := tabstack.AgentResearchParams{Query: in.Query}
		if in.Mode != "" {
			body.Mode = tabstack.AgentResearchParamsMode(in.Mode)
		}
		if in.Nocache {
			body.Nocache = param.NewOpt(true)
		}
		if in.FetchTimeout > 0 {
			body.FetchTimeout = param.NewOpt(in.FetchTimeout)
		}

		emitter := newProgressEmitter(ctx, req)
		stream := c.Agent.ResearchStreaming(ctx, body)
		defer func() { _ = stream.Close() }()

		var report string
		for stream.Next() {
			if err := ctx.Err(); err != nil {
				return toolError(err), nil, nil
			}
			ev := stream.Current()
			if msg := researchProgress(ev); msg != "" {
				emitter.emit(msg)
			}
			switch v := ev.AsAny().(type) {
			case tabstack.ResearchEventComplete:
				report = v.Data.Report
			case tabstack.ResearchEventError:
				return toolError(fmt.Errorf("research: %s", v.Data.Error.Message)), nil, nil
			}
		}
		if err := stream.Err(); err != nil {
			return toolError(fmt.Errorf("research stream: %w", err)), nil, nil
		}
		if report == "" {
			return toolError(fmt.Errorf("research stream ended without a complete event")), nil, nil
		}
		return &sdk.CallToolResult{
			Content: []sdk.Content{&sdk.TextContent{Text: report}},
		}, nil, nil
	})
}
