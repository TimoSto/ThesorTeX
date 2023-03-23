package example_project

import _ "embed"

//go:embed example_de.zip
var ExampleDE []byte

//go:embed documentation_de.pdf
var PdfExampleDE []byte

//go:embed example_en.zip
var ExampleEN []byte

//go:embed documentation_en.pdf
var PdfExampleEN []byte

func GetExampleZip(lang string) []byte {
	if lang == "de" {
		return ExampleDE
	}
	return ExampleEN
}

func GetExamplePDF(lang string) []byte {
	if lang == "de" {
		return PdfExampleDE
	}
	return PdfExampleEN
}
