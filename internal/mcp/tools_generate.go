package mcp

import (
	"context"
	"fmt"

	sdk "github.com/modelcontextprotocol/go-sdk/mcp"
	tabstack "github.com/stainless-sdks/tabstack-go"
	"github.com/stainless-sdks/tabstack-go/packages/param"
)

// GenerateJsonInput is the input for tabstack_generate_json.
type GenerateJsonInput struct {
	URL          string         `json:"url" jsonschema:"URL to fetch content from"`
	JsonSchema   map[string]any `json:"json_schema" jsonschema:"JSON Schema describing the structure of the transformed output"`
	Instructions string         `json:"instructions" jsonschema:"transformation instructions (max 20000 chars)"`
	Nocache      bool           `json:"nocache,omitempty" jsonschema:"bypass server-side cache"`
	Effort       string         `json:"effort,omitempty" jsonschema:"min, standard, or max"`
	Geo          string         `json:"geo,omitempty" jsonschema:"ISO 3166-1 alpha-2 country code"`
}

func registerGenerateJson(s *sdk.Server, c *tabstack.Client) {
	sdk.AddTool(s, &sdk.Tool{
		Name: "tabstack_generate_json",
		Description: "Fetch a URL and AI-transform its content per your instructions, returning JSON " +
			"that conforms to a schema you provide. Use when you need a synthesized or summarized version " +
			"of a page in a structured shape (e.g. a one-paragraph summary, a sentiment classification, " +
			"a key-points list).",
	}, func(ctx context.Context, _ *sdk.CallToolRequest, in *GenerateJsonInput) (*sdk.CallToolResult, any, error) {
		body := tabstack.GenerateJsonParams{
			URL:          in.URL,
			JsonSchema:   in.JsonSchema,
			Instructions: in.Instructions,
		}
		if in.Nocache {
			body.Nocache = param.NewOpt(true)
		}
		if in.Effort != "" {
			body.Effort = tabstack.GenerateJsonParamsEffort(in.Effort)
		}
		if in.Geo != "" {
			body.GeoTarget = tabstack.GenerateJsonParamsGeoTarget{Country: param.NewOpt(in.Geo)}
		}
		resp, err := c.Generate.Json(ctx, body)
		if err != nil {
			return toolError(fmt.Errorf("generate json: %w", err)), nil, nil
		}
		return jsonResult(resp), nil, nil
	})
}
