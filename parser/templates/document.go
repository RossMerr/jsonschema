package templates

import (
	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser"
)

var _ parser.Root = (*Document)(nil)

type Document struct {
	ID      string
	Package string
	globals map[string]parser.Component
	Imports []string
}

func NewDocument(packagename string, schema *jsonschema.Schema) *Document {
	schema.SetParent("", nil)

	document := &Document{
		ID:      schema.ID.String(),
		globals: map[string]parser.Component{},
		Package: packagename,
	}

	return document
}

func (s *Document) AddImport(value string) {
	s.Imports = append(s.Imports, value)
	s.Imports = unique(s.Imports)
}

func (s *Document) Globals() map[string]parser.Component {
	return s.globals
}

func (s *Document) Name() string {
	return EmptyString
}

func (s *Document) WithPackageName(packagename string) {
	s.Package = packagename
}

func unique(slice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

const DocumentTemplate = `
// Code generated by jsonschema. DO NOT EDIT.

{{- if .Package}}
package {{.Package}}
{{else}}
package main
{{- end}}

{{range $key, $import := .Imports -}}
import "{{$import}}"
{{end}}

{{ if .ID -}}
// ID: {{.ID}}
{{ end }}

{{range $key, $global := .Globals -}}
	{{- template "kind" $global -}}
{{end}}
`
