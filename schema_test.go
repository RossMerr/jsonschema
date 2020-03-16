package jsonschema_test

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/interpreter"
	"github.com/RossMerr/jsonschema/parser"
)

func TestSchemas_Generate(t *testing.T) {
	type fields struct {
		documents map[jsonschema.ID]*jsonschema.Schema
		config    *jsonschema.Config
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{

		{
			name: "Basic",
			fields: fields{
				documents: map[jsonschema.ID]*jsonschema.Schema{
					"basicBasic": loadRawSchema("samples/basicBasic.json"),
				},
				config: &jsonschema.Config{
					Packagename: "main",
					Output: "output/",
				},
			},
		},
		{
			name: "Nesting data structures",
			fields: fields{
				documents: map[jsonschema.ID]*jsonschema.Schema{
					"productNesting": loadRawSchema("samples/productNesting.json"),
				},
				config: &jsonschema.Config{
					Packagename: "main",
					Output:      "output/",
				},
			},
		},
		{
			name: "References outside the schema",
			fields: fields{
				documents: map[jsonschema.ID]*jsonschema.Schema{
					"https://example.com/geographical-location.schema.json": loadRawSchema("samples/geographical-location.schema.json"),
					"http://example.com/product.schema.json": loadRawSchema("samples/product.schema.json"),

				},
				config: &jsonschema.Config{
					Packagename: "main",
					Output: "output/",
				},
			},
		},
		{
			name: "Oneof",
			fields: fields{
				documents: map[jsonschema.ID]*jsonschema.Schema{
					"http://example.com/entry-schema": loadRawSchema("samples/entry-schema.json"),
				},
				config: &jsonschema.Config{
					Packagename: "main",
					Output:      "output/",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			p := parser.NewParser(context.Background(), tt.fields.config.Packagename)
			parse := p.Parse(tt.fields.documents)
			interpret := interpreter.NewInterpretDefaults(parse)

			if err := interpret.ToFile(tt.fields.config.Output); (err != nil) != tt.wantErr {
				t.Errorf("Schemas.Generate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func loadRawSchema(filename string) *jsonschema.Schema {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var doc jsonschema.Schema
	err = json.Unmarshal(data, &doc)
	if err != nil {
		panic(err)
	}

	return &doc
}
