package handlers

import (
	"fmt"
	"strconv"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser"
	"github.com/RossMerr/jsonschema/parser/templates"
)

func HandleAnyOf(ctx *parser.SchemaContext, doc parser.Root, name string, schema *jsonschema.Schema) (parser.Component, error) {

	methodSignature := parser.NewMethodSignature(name)

	types := make([]string, 0)
	for i, subschema := range schema.AnyOf {
		if subschema.Ref.IsNotEmpty() {
			receiver := subschema.Ref.ToKey()
			types = append(types, receiver)
			ctx.RegisterMethodSignature(receiver, methodSignature)
			continue
		}

		structname := name + strconv.Itoa(i)
		t, err := ctx.Process(doc, structname, subschema)

		if err != nil {
			return nil, err
		}

		_, contains := doc.Globals()[structname]
		if contains {
			return nil, fmt.Errorf("handleanyof: anyOf, global keys need to be unique found %v more than once", structname)
		}

		doc.Globals()[structname] = templates.NewType(schema.Description, t)
		types = append(types, structname)
		ctx.RegisterMethodSignature(structname, methodSignature)
	}

	doc.AddImport("encoding/json")
	doc.Globals()[name] = templates.NewInterface(name).WithMethodSignature(methodSignature)
	r := templates.NewReference(name, "", parser.NewType(name, parser.Array), types...)
	return &templates.AnyOf{r}, nil
}