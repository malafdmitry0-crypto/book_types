package main

import (
	"bytes"
	_ "embed"
	"go/format"
	"os"
	"text/template"
)

//go:embed algorithms.tmpl
var source string

func main() {
	tmpl, err := template.New("algorithms").Parse(source)
	if err != nil {
		panic(err)
	}
	var out bytes.Buffer
	if err := tmpl.Execute(&out, struct{ Type string }{"fraction.Fraction"}); err != nil {
		panic(err)
	}
	code, err := format.Source(out.Bytes())
	if err != nil {
		panic(err)
	}
	if err := os.WriteFile("generated.go", code, 0644); err != nil {
		panic(err)
	}
}
