package mcp

import (
	"context"
	"fmt"

	sdk "github.com/modelcontextprotocol/go-sdk/mcp"
	tabstack "github.com/stainless-sdks/tabstack-go"
	"github.com/stainless-sdks/tabstack-go/packages/param"
)

// ExtractMarkdownInput is the JSON-Schema-derived input for tabstack_extract_markdown.
type ExtractMarkdownInput struct {
	URL      string `json:"url" jsonschema:"URL to fetch and convert to clean markdown"`
	Metadata bool   `json:"metadata,omitempty" jsonschema:"return metadata as a structured field instead of YAML frontmatter inside content"`
	Nocache  bool   `json:"nocache,omitempty" jsonschema:"bypass server-side cache"`
	Effort   string `json:"effort,omitempty" jsonschema:"speed/capability tradeoff: min, standard, or max"`
	Geo      string `json:"geo,omitempty" jsonschema:"ISO 3166-1 alpha-2 country code (e.g. US, GB, JP)"`
}

func registerExtractMarkdown(s *sdk.Server, c *tabstack.Client) {
	sdk.AddTool(s, &sdk.Tool{
		Name: "tabstack_extract_markdown",
		Description: "Fetch a URL and return its main content as clean markdown. " +
			"Use when you need the readable text of a web page without HTML or layout boilerplate.",
	}, func(ctx context.Context, _ *sdk.CallToolRequest, in *ExtractMarkdownInput) (*sdk.CallToolResult, any, error) {
		body := tabstack.ExtractMarkdownParams{URL: in.URL}
		if in.Metadata {
			body.Metadata = param.NewOpt(true)
		}
		if in.Nocache {
			body.Nocache = param.NewOpt(true)
		}
		if in.Effort != "" {
			body.Effort = tabstack.ExtractMarkdownParamsEffort(in.Effort)
		}
		if in.Geo != "" {
			body.GeoTarget = tabstack.ExtractMarkdownParamsGeoTarget{Country: param.NewOpt(in.Geo)}
		}
		resp, err := c.Extract.Markdown(ctx, body)
		if err != nil {
			return toolError(fmt.Errorf("extract markdown: %w", err)), nil, nil
		}
		return jsonResult(resp), nil, nil
	})
}
