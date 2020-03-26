package parser

import (
	"strconv"

	"github.com/RossMerr/jsonschema"
)

type InterfaceReference struct {
	Type string
	name     string
	FieldTag string
}

func NewInterfaceReference(ctx *SchemaContext, name *Name, fieldTag string, subschemas []*jsonschema.Schema) *InterfaceReference {
	parent := ctx.Parent()

	filename := parent.ID.Filename()
	typename := jsonschema.Structname(filename) + name.Fieldname()

	for i, subschema := range subschemas {
		if !subschema.Ref.IsEmpty() {
			ctx.AddMethods(subschema.Ref.Fieldname(), typename)
			continue
		}
		structname := typename + strconv.Itoa(i)
		t :=  SchemaToType(ctx, NewName(structname), subschema, false)
		ctx.Globals[structname] = PrefixType(t, typename)
		ctx.AddMethods(structname, typename)

	}

	ctx.Globals[name.Fieldname()] = NewInterface(typename)

	return &InterfaceReference{
		Type: typename,
		name: name.Fieldname(),
		FieldTag: fieldTag,
	}
}

func (s *InterfaceReference) Comment() string {
	return jsonschema.EmptyString
}

func (s *InterfaceReference) Name() string {
	return s.name
}

const InterfaceReferenceTemplate = `
{{- define "interfacereference" -}}
{{ .Name}} {{ .Type}} {{ .FieldTag }}
{{- end -}}
`