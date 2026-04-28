package mcp

import (
	"fmt"
	"reflect"

	"github.com/google/jsonschema-go/jsonschema"
)

// inputSchemaWithEnums builds an input schema for T, then sets `enum` on the
// named JSON properties so the LLM sees the allowed values. The jsonschema-go
// tag-based inference only supports descriptions; enum constraints have to be
// patched in programmatically.
func inputSchemaWithEnums[T any](enums map[string][]string) *jsonschema.Schema {
	schema, err := jsonschema.ForType(reflect.TypeFor[T](), &jsonschema.ForOptions{})
	if err != nil {
		panic(fmt.Errorf("inputSchemaWithEnums[%T]: %w", *new(T), err))
	}
	for prop, values := range enums {
		ps, ok := schema.Properties[prop]
		if !ok {
			panic(fmt.Errorf("inputSchemaWithEnums[%T]: property %q not in schema", *new(T), prop))
		}
		ps.Enum = make([]any, len(values))
		for i, v := range values {
			ps.Enum[i] = v
		}
	}
	return schema
}
