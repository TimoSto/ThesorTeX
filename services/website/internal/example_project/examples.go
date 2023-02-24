package example_project

import _ "embed"

//go:embed example_de.zip
var ExampleDE []byte

//go:embed example_en.zip
var ExampleEN []byte

func GetExampleZip(lang string) []byte {
	if lang == "de" {
		return ExampleDE
	}
	return ExampleEN
}
