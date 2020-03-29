package interpreter

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/RossMerr/jsonschema/parser"
	"github.com/gookit/color"
	log "github.com/sirupsen/logrus"
)

type Interpret interface {
	ToFile(output string) ([]string, error)
}

type interpret struct {
	root           *parser.Parse
	templateStruct Template
}

func NewInterpret(root *parser.Parse, templateStruct Template) Interpret {
	return &interpret{
		root:           root,
		templateStruct: templateStruct,
	}
}

func NewInterpretDefaults(root *parser.Parse) (Interpret, error) {
	templates, err := parser.Template()
	if err != nil {
		return nil, err
	}
	return NewInterpret(root, templates), nil
}

func (s *interpret) ToFile(output string) ([]string, error) {
	files := []string{}
	green := color.FgGreen.Render

	for _, obj := range s.root.Structs {
		filename := path.Join(output, obj.Filename+".go")

		_, err := os.Stat(filename)
		if !os.IsNotExist(err) {
			err = os.Remove(filename)
			if err != nil {
				return files, err
			}
		}

		file, err := os.Create(filename)
		if err != nil {
			return files, err
		}
		files = append(files, filename)

		err = s.templateStruct.Execute(file, obj)
		if err != nil {
			return files, err
		}

		log.Infof("Create file %v", filename)

		cmd := exec.Command("gofmt", "-w", filename)
		if err := cmd.Run(); err != nil {
			return files, err
		}
	}

	fmt.Printf(green("✓") +" Create %v files\n", len(s.root.Structs))

	return files, nil
}
