package cmd

import (
	"fmt"
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
	validStreamOutputs = []string{"json", "pretty"}
)

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
