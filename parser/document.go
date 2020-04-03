package parser

import (
	"github.com/RossMerr/jsonschema"
)

type Process func(name string, schema *jsonschema.Schema) (Types, error)
type HandleSchemaFunc func(*SchemaContext, *Document, string, *jsonschema.Schema) (Types, error)

type Document struct {
	ID         string
	Package    string
	Globals    map[string]Types
	Filename   string
	rootSchema *jsonschema.Schema
}

func (s *Document) Root() *jsonschema.Schema {
	return s.rootSchema
}

const DocumentTemplate = `
// Code generated by jsonschema. DO NOT EDIT.

{{- if .Package}}
package {{.Package}}
{{else}}
package main
{{- end}}

{{ if .ID -}}
// ID: {{.ID}}
{{ end }}

{{range $key, $global := .Globals -}}
	{{- template "kind" $global -}}
{{end}}
`
