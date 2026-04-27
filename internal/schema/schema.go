// Package schema loads JSON Schema documents for Tabstack extract/generate
// requests. Schemas are passed through verbatim — no validation here.
package schema

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Load reads a JSON document from the given path and returns it as a generic
// value (typically map[string]any) suitable for the Tabstack SDK's JsonSchema
// field. The path "-" reads from stdin.
func Load(path string) (any, error) {
	var r io.Reader
	if path == "-" {
		r = os.Stdin
	} else {
		f, err := os.Open(path)
		if err != nil {
			return nil, fmt.Errorf("open schema %q: %w", path, err)
		}
		defer func() { _ = f.Close() }()
		r = f
	}

	var schema any
	if err := json.NewDecoder(r).Decode(&schema); err != nil {
		return nil, fmt.Errorf("parse schema %q as JSON: %w", path, err)
	}
	return schema, nil
}
