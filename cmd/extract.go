package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/lmorchard/tabstack-go-cli/internal/client"
	"github.com/lmorchard/tabstack-go-cli/internal/schema"
	"github.com/spf13/cobra"
	tabstack "github.com/stainless-sdks/tabstack-go"
	"github.com/stainless-sdks/tabstack-go/packages/param"
)

var extractCmd = &cobra.Command{
	Use:   "extract",
	Short: "Extract content from a URL",
	Long:  `Extract markdown or structured JSON from a fetched URL.`,
}

var (
	extractMarkdownMetadata bool
	extractMarkdownNocache  bool
	extractMarkdownEffort   string
	extractMarkdownGeo      string
)

var extractMarkdownCmd = &cobra.Command{
	Use:   "markdown <url>",
	Short: "Fetch a URL and convert to clean markdown",
	Args:  cobra.ExactArgs(1),
	RunE:  runExtractMarkdown,
}

func runExtractMarkdown(_ *cobra.Command, args []string) error {
	if err := validateEnum("effort", extractMarkdownEffort, validEfforts); err != nil {
		return err
	}
	c, err := client.New(GetConfig())
	if err != nil {
		return err
	}

	body := tabstack.ExtractMarkdownParams{
		URL: args[0],
	}
	if extractMarkdownMetadata {
		body.Metadata = param.NewOpt(true)
	}
	if extractMarkdownNocache {
		body.Nocache = param.NewOpt(true)
	}
	if extractMarkdownEffort != "" {
		body.Effort = tabstack.ExtractMarkdownParamsEffort(extractMarkdownEffort)
	}
	if extractMarkdownGeo != "" {
		body.GeoTarget = tabstack.ExtractMarkdownParamsGeoTarget{
			Country: param.NewOpt(extractMarkdownGeo),
		}
	}

	resp, err := c.Extract.Markdown(context.Background(), body)
	if err != nil {
		return fmt.Errorf("extract markdown: %w", err)
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(resp)
}

func init() {
	extractMarkdownCmd.Flags().BoolVar(&extractMarkdownMetadata, "metadata", false, "include extracted metadata")
	extractMarkdownCmd.Flags().BoolVar(&extractMarkdownNocache, "nocache", false, "bypass cache")
	extractMarkdownCmd.Flags().StringVar(&extractMarkdownEffort, "effort", "", "effort level: min, standard, or max")
	extractMarkdownCmd.Flags().StringVar(&extractMarkdownGeo, "geo", "", "ISO 3166-1 alpha-2 country code (e.g. US, GB)")

	extractCmd.AddCommand(extractMarkdownCmd)
	rootCmd.AddCommand(extractCmd)
	initExtractJson()
}

var (
	extractJsonSchemaPath string
	extractJsonNocache    bool
	extractJsonEffort     string
	extractJsonGeo        string
)

var extractJsonCmd = &cobra.Command{
	Use:   "json <url>",
	Short: "Fetch a URL and extract structured data per a JSON Schema",
	Args:  cobra.ExactArgs(1),
	RunE:  runExtractJson,
}

func runExtractJson(_ *cobra.Command, args []string) error {
	if extractJsonSchemaPath == "" {
		return fmt.Errorf("--schema is required")
	}
	if err := validateEnum("effort", extractJsonEffort, validEfforts); err != nil {
		return err
	}
	sch, err := schema.Load(extractJsonSchemaPath)
	if err != nil {
		return err
	}

	c, err := client.New(GetConfig())
	if err != nil {
		return err
	}

	body := tabstack.ExtractJsonParams{
		URL:        args[0],
		JsonSchema: sch,
	}
	if extractJsonNocache {
		body.Nocache = param.NewOpt(true)
	}
	if extractJsonEffort != "" {
		body.Effort = tabstack.ExtractJsonParamsEffort(extractJsonEffort)
	}
	if extractJsonGeo != "" {
		body.GeoTarget = tabstack.ExtractJsonParamsGeoTarget{
			Country: param.NewOpt(extractJsonGeo),
		}
	}

	resp, err := c.Extract.Json(context.Background(), body)
	if err != nil {
		return fmt.Errorf("extract json: %w", err)
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(resp)
}

func initExtractJson() {
	extractJsonCmd.Flags().StringVar(&extractJsonSchemaPath, "schema", "", `path to JSON Schema file (use "-" for stdin)`)
	extractJsonCmd.Flags().BoolVar(&extractJsonNocache, "nocache", false, "bypass cache")
	extractJsonCmd.Flags().StringVar(&extractJsonEffort, "effort", "", "effort level: min, standard, or max")
	extractJsonCmd.Flags().StringVar(&extractJsonGeo, "geo", "", "ISO 3166-1 alpha-2 country code")

	_ = extractJsonCmd.MarkFlagRequired("schema")
	extractCmd.AddCommand(extractJsonCmd)
}
