package types

import (
	"github.com/RossMerr/jsonschema/parser"
)

var _ parser.Types = (*List)(nil)

type List struct {
	Items []parser.Types
}

func (s *List) WithReference(ref bool) parser.Types {
	return s
}

func (s *List) WithFieldTag(tags string) parser.Types {
	return s
}

func (s *List) Comment() string {
	return ""
}

func (s *List) Name() string {
	return ""
}
