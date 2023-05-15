package querycost

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Output interface {
	DisplayResult(res FormattedResult)
}
type StandardOutput struct{}
type CSVOutput struct {
	file string
}

func (s StandardOutput) DisplayResult(res FormattedResult) {
	for _, row := range res.Value {
		fmt.Println(strings.Join(row, ","))
	}
}
func (s CSVOutput) DisplayResult(res FormattedResult) {
	f, err := os.Create(s.file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for _, row := range res.Value {
		f.WriteString(strings.Join(row, ","))
		f.WriteString("\n")
	}
}

func NewCSVOutput(_output string) CSVOutput {
	parent := filepath.Dir(_output)
	if err := os.MkdirAll(parent, os.ModePerm); err != nil {
		panic(err)
	}
	return CSVOutput{file: _output}
}

type ArrayOutput struct {
	Array *[][]string
}

func NewArrayOutput(output *[][]string) ArrayOutput {
	return ArrayOutput{Array: output}
}

func (a ArrayOutput) DisplayResult(res FormattedResult) {
	a.Array = &res.Value
}
