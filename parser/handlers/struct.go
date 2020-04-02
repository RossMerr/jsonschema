package handlers

import (
	"strings"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser/document"
	"github.com/RossMerr/jsonschema/parser/tags"
	"github.com/RossMerr/jsonschema/parser/tags/json"
	"github.com/RossMerr/jsonschema/parser/tags/validate"
	"github.com/RossMerr/jsonschema/parser/types"
)

func HandleObject(doc *document.Document, name string, schema *jsonschema.Schema) (document.Types, error) {

	fields := []document.Types{}
	for key, propertie := range schema.Properties {
		s, err := doc.Process(key, propertie)
		if err != nil {
			return nil, err
		}

		tags := tags.NewFieldTag([]tags.StructTag{json.NewJSONTags(), validate.NewValidateTags()})
		fieldTag := tags.ToFieldTag(key, propertie, schema.Required)

		ref := !jsonschema.Contains(schema.Required, strings.ToLower(key))

		s.WithFieldTag(fieldTag).WithReference(ref)

		if _, ok  := s.(*types.Root); ok {
			continue
		}

		fields = append(fields, s)
	}

	for key, def := range schema.AllDefinitions() {
		s, err := HandleRoot(doc, key, def)
		if err != nil {
			return nil, err
		}

		if _, contains := doc.Globals[key]; !contains {
			doc.Globals[key] = s
		}
	}

	return types.NewStruct(name, schema.Description, fields), nil
}
