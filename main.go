package main

import (
	"errors"
	"image"
	"os"
	"path/filepath"
	"strings"

	"github.com/actions-go/toolkit/core"
	svgmap "github.com/owulveryck/wardleyToGo/encoding/svg"
	"github.com/owulveryck/wardleyToGo/parser/owm"
)

func runMain() error {
	input, ok := core.GetInput("OWMFILE")
	if !ok {
		return errors.New("invalid input")
	}
	f, err := os.Open(input)
	if err != nil {
		return err
	}
	defer f.Close()
	p := owm.NewParser(f)
	m, err := p.Parse() // the map
	if err != nil {
		return err
	}
	output := strings.Trim(input, filepath.Ext(input)) + ".svg"
	outputF, err := os.Create(output)
	if err != nil {
		return err
	}
	defer outputF.Close()
	e, err := svgmap.NewEncoder(outputF, image.Rect(0, 0, 1100, 900), image.Rect(30, 50, 1070, 850))
	if err != nil {
		return err
	}
	defer e.Close()
	style := svgmap.NewWardleyStyle(svgmap.DefaultEvolution)
	e.Init(style)
	err = e.Encode(m)
	if err != nil {
		return err
	}
	core.SetOutput("svgfile", "output")
	return nil
}

func main() {
	err := runMain()
	if err != nil {
		core.Error(err.Error())
	}
}
