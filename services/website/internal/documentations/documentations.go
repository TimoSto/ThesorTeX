package documentations

import (
	"embed"
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
)

//go:embed docs/thesisTemplate/de/parsed.json
var templateDocJSON_DE []byte

//go:embed docs/thesisTemplate/en/parsed.json
var templateDocJSON_EN []byte

//go:embed docs/cvTemplate/de/parsed.json
var cvTemplateDocJSON_DE []byte

//go:embed docs/cvTemplate/en/parsed.json
var cvTemplateDocJSON_EN []byte

//go:embed documentation_de.pdf
var docPDF_DE []byte

//go:embed documentation_en.pdf
var docPDF_EN []byte

//go:embed docsImages
var docsImages embed.FS

var (
	ErrNotFound = errors.New("documentation not found")
)

var docsMap = map[string][]byte{
	"thesisTemplate_JSON_DE": templateDocJSON_DE,
	"thesisTemplate_JSON_EN": templateDocJSON_EN,
	"cvTemplate_JSON_DE":     cvTemplateDocJSON_DE,
	"cvTemplate_JSON_EN":     cvTemplateDocJSON_EN,
	"documantation_PDF_DE":   docPDF_DE,
	"documantation_PDF_EN":   docPDF_EN,
}

type ThesorTexDocs struct {
	ThesisTemplate string
	CVTemplate     string
}

func GetJsonDocs(lang string) ([]byte, error) {
	lang = strings.ToUpper(lang)
	if lang != "DE" {
		lang = "EN"
	}

	var docs ThesorTexDocs
	docs.ThesisTemplate = string(docsMap[fmt.Sprintf("thesisTemplate_JSON_%s", lang)])
	docs.CVTemplate = string(docsMap[fmt.Sprintf("cvTemplate_JSON_%s", lang)])

	return json.Marshal(docs)
}

func GetPDFDocs(lang string) ([]byte, error) {
	lang = strings.ToUpper(lang)
	if lang != "DE" {
		lang = "EN"
	}

	return docPDF_DE, nil
}

func GetDocsImage(path string) ([]byte, error) {
	fs.WalkDir(docsImages, ".", func(path string, d fs.DirEntry, err error) error {
		fmt.Println(path)
		return nil
	})
	path = filepath.Join("docsImages", path)
	return docsImages.ReadFile(path)
}
