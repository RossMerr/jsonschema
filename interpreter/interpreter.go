package interpreter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/RossMerr/jsonschema"
	"github.com/RossMerr/jsonschema/parser"
	"github.com/RossMerr/jsonschema/parser/types"
	"github.com/gookit/color"

	log "github.com/sirupsen/logrus"
)

type Interpreter struct {
	parser parser.Parser
}

func NewInterpreter(parser parser.Parser) *Interpreter {
	return &Interpreter{
		parser: parser,
	}
}

func NewInterpreterDefaults(packagename string) *Interpreter {
	p := parser.NewParser(packagename)
	p.HandlerFunc(parser.Boolean, types.HandleBoolean)
	p.HandlerFunc(parser.OneOf, types.HandleOneOf)
	p.HandlerFunc(parser.AnyOf, types.HandleAnyOf)
	p.HandlerFunc(parser.AllOf, types.HandleAllOf)
	p.HandlerFunc(parser.Enum, types.HandleEnum)
	p.HandlerFunc(parser.Array, types.HandleArray)
	p.HandlerFunc(parser.Reference, types.HandleReference)
	p.HandlerFunc(parser.Object, types.HandleObject)
	p.HandlerFunc(parser.Number, types.HandleNumber)
	p.HandlerFunc(parser.Interger, types.HandleInteger)
	p.HandlerFunc(parser.String, types.HandleString)
	p.HandlerFunc(parser.RootObject, types.HandleRoot)

	return NewInterpreter(p)
}

func (s *Interpreter) Interpret(files []string) (Interpret, error) {
	schemas := map[jsonschema.ID]jsonschema.JsonSchema{}
	references := map[jsonschema.ID]jsonschema.JsonSchema{}
	green := color.FgCyan.Render
	red := color.FgRed.Render

	rawFiles := map[string][]byte{}
	for _, filename := range files {

		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Printf(red("🗴") + "Reading file\n")
			return nil, err
		}

		log.Infof("Found file %v", filename)
		rawFiles[filename] = data
	}

	fmt.Printf(green("✓")+" Found %v files\n", len(rawFiles))

	i := 0
	for _, data := range rawFiles {
		var schema jsonschema.Schema
		err := json.Unmarshal(data, &schema)
		if err != nil {
			fmt.Printf(red("🗴") + "Unmarshalling\n")
			return nil, err
		} else {
			i++
		}
		log.Infof("Found schema %v", schema.ID)

		schemas[schema.ID] = &schema
	}

	fmt.Printf(green("✓")+" Unmarshalled %v schemas\n", i)

	for _, data := range rawFiles {

		refs := jsonschema.ResolveIDs(data)

		for k, v := range refs {
			if _, ok := references[k]; !ok {
				references[k] = v
			} else {
				fmt.Printf(red("🗴") + "References\n")

				return nil, fmt.Errorf("Reference keys need to be unique found %v more than once", k)
			}
			log.Infof("Found reference %v", k)
		}
	}

	fmt.Printf(green("✓")+" Found %v references\n", len(references))

	root, err := s.parser.Parse(schemas, references)
	if err != nil {
		return nil, err
	}
	return NewInterpretDefaults(root)
}
