package templates

import (
	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser"
)

var _ parser.Types = (*Struct)(nil)

type Struct struct {
	comment  string
	name     string
	Fields   []parser.Types
	fieldTag string
}

func NewStruct(name, comment string, fields []parser.Types) *Struct {
	return &Struct{
		comment: comment,
		name:    jsonschema.ToTypename(name),
		Fields:  fields,
	}
}

func (s *Struct) WithMethods(methods ...*parser.Method) parser.Types {
	return s
}

func (s *Struct) WithReference(ref bool) parser.Types {
	return s
}

func (s *Struct) WithFieldTag(tags string) parser.Types {
	s.fieldTag = tags
	return s
}

func (s *Struct) FieldTag() string {
	return s.fieldTag
}

func (s *Struct) Comment() string {
	return s.comment
}

func (s *Struct) Name() string {
	return s.name
}

func (s *Struct) IsNotEmpty() bool {
	return len(s.Fields) > 0
}

const StructTemplate = `
{{- define "struct" -}}
{{ .Name }} struct {
{{range $key, $propertie := .Fields -}}
	{{template "kind" $propertie }}
{{end -}}
} {{ .FieldTag }}
{{- end -}}`
