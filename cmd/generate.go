package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/lmorchard/tabstack-go-cli/internal/client"
	"github.com/lmorchard/tabstack-go-cli/internal/schema"
	"github.com/spf13/cobra"
	tabstack "github.com/stainless-sdks/tabstack-go"
	"github.com/stainless-sdks/tabstack-go/packages/param"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate AI-transformed content from a URL",
}

var (
	generateJsonSchemaPath       string
	generateJsonInstructions     string
	generateJsonInstructionsFile string
	generateJsonNocache          bool
	generateJsonEffort           string
	generateJsonGeo              string
)

var generateJsonCmd = &cobra.Command{
	Use:   "json <url>",
	Short: "Fetch a URL and AI-transform it into JSON per a schema and instructions",
	Args:  cobra.ExactArgs(1),
	RunE:  runGenerateJson,
}

func runGenerateJson(_ *cobra.Command, args []string) error {
	if err := validateEnum("effort", generateJsonEffort, validEfforts); err != nil {
		return err
	}
	instructions, err := resolveInstructions(generateJsonInstructions, generateJsonInstructionsFile)
	if err != nil {
		return err
	}
	sch, err := schema.Load(generateJsonSchemaPath)
	if err != nil {
		return err
	}

	c, err := client.New(GetConfig())
	if err != nil {
		return err
	}

	body := tabstack.GenerateJsonParams{
		URL:          args[0],
		Instructions: instructions,
		JsonSchema:   sch,
	}
	if generateJsonNocache {
		body.Nocache = param.NewOpt(true)
	}
	if generateJsonEffort != "" {
		body.Effort = tabstack.GenerateJsonParamsEffort(generateJsonEffort)
	}
	if generateJsonGeo != "" {
		body.GeoTarget = tabstack.GenerateJsonParamsGeoTarget{
			Country: param.NewOpt(generateJsonGeo),
		}
	}

	resp, err := c.Generate.Json(context.Background(), body)
	if err != nil {
		return fmt.Errorf("generate json: %w", err)
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(resp)
}

// resolveInstructions returns the instructions string from either the inline
// flag or a file path. Exactly one must be set.
func resolveInstructions(inline, path string) (string, error) {
	if inline == "" && path == "" {
		return "", fmt.Errorf("provide --instructions or --instructions-file")
	}
	if inline != "" && path != "" {
		return "", fmt.Errorf("--instructions and --instructions-file are mutually exclusive")
	}
	if inline != "" {
		return inline, nil
	}
	b, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("read instructions: %w", err)
	}
	return strings.TrimSpace(string(b)), nil
}

func init() {
	generateJsonCmd.Flags().StringVar(&generateJsonSchemaPath, "schema", "", `path to JSON Schema file (use "-" for stdin)`)
	generateJsonCmd.Flags().StringVar(&generateJsonInstructions, "instructions", "", "transformation instructions (max 20000 chars)")
	generateJsonCmd.Flags().StringVar(&generateJsonInstructionsFile, "instructions-file", "", "path to a file containing the instructions")
	generateJsonCmd.Flags().BoolVar(&generateJsonNocache, "nocache", false, "bypass cache")
	generateJsonCmd.Flags().StringVar(&generateJsonEffort, "effort", "", "effort level: min, standard, or max")
	generateJsonCmd.Flags().StringVar(&generateJsonGeo, "geo", "", "ISO 3166-1 alpha-2 country code")

	_ = generateJsonCmd.MarkFlagRequired("schema")

	generateCmd.AddCommand(generateJsonCmd)
	rootCmd.AddCommand(generateCmd)
}
