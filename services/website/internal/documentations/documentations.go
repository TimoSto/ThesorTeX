package documentations

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

//go:embed docs/thesisTemplate/de/parsed.json
var templateDocJSON_DE []byte

//go:embed docs/thesisTemplate/en/parsed.json
var templateDocJSON_EN []byte

var (
	ErrNotFound = errors.New("documentation not found")
)

var docsMap = map[string][]byte{
	"thesisTemplate_JSON_DE": templateDocJSON_DE,
	"thesisTemplate_JSON_EN": templateDocJSON_EN,
}

type ThesorTexDocs struct {
	ThesisTemplate string
}

func GetJsonDocs(lang string) ([]byte, error) {
	lang = strings.ToUpper(lang)
	if lang != "DE" {
		lang = "EN"
	}

	var docs ThesorTexDocs
	docs.ThesisTemplate = string(docsMap[fmt.Sprintf("thesisTemplate_JSON_%s", lang)])

	return json.Marshal(docs)
}
