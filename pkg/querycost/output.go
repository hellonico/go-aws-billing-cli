package querycost

import (
	"fmt"
	"log"
	"os"
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
