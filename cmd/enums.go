package cmd

import (
	"fmt"
	"os"
	"strings"
)

// Valid values for the API's enum-typed flags. Hardcoded as plain strings
// rather than referencing the SDK's typed constants because the three SDK
// effort types (ExtractMarkdownParamsEffort, ExtractJsonParamsEffort,
// GenerateJsonParamsEffort) are distinct named types that today happen to
// share the same valid set. If the API ever differentiates them, split these.
var (
	validEfforts       = []string{"min", "standard", "max"}
	validResearchModes = []string{"fast", "balanced"}
	validStreamOutputs = []string{"auto", "json", "pretty"}
	validColorModes    = []string{"auto", "always", "never"}
	validMCPTransports = []string{"stdio", "http"}
)

// resolveStreamOutput maps the --output flag's "auto|json|pretty" choice into
// the concrete renderer to use. "auto" picks pretty when stdout is a TTY and
// json otherwise — so interactive use gets human-readable output by default
// while shell pipelines (`tabstack research foo | jq ...`) keep getting JSON.
func resolveStreamOutput(mode string) string {
	switch mode {
	case "json", "pretty":
		return mode
	default: // "auto" or empty
		if isTerminal(os.Stdout) {
			return "pretty"
		}
		return "json"
	}
}

// resolveStreamColor maps the --color flag's "auto|always|never" choice into
// the bool the pretty renderer takes. "auto" enables color when stdout is a
// TTY AND NO_COLOR is unset (the de-facto opt-out — see https://no-color.org).
func resolveStreamColor(mode string) bool {
	switch mode {
	case "always":
		return true
	case "never":
		return false
	default: // "auto" or empty
		if os.Getenv("NO_COLOR") != "" {
			return false
		}
		return isTerminal(os.Stdout)
	}
}

// validateEnum returns an error when value is non-empty and not in valid.
// An empty value is permitted — the API uses its own default.
func validateEnum(flagName, value string, valid []string) error {
	if value == "" {
		return nil
	}
	for _, v := range valid {
		if value == v {
			return nil
		}
	}
	return fmt.Errorf("invalid --%s %q: must be one of %s", flagName, value, strings.Join(valid, ", "))
}
