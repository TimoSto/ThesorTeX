package documentations

import (
	_ "embed"
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

func GetDocumentation(topic string, format string, lang string) ([]byte, error) {
	lang = strings.ToUpper(lang)
	if lang != "DE" {
		lang = "EN"
	}

	doc := docsMap[fmt.Sprintf("%s_%s_%s", topic, format, lang)]

	if len(doc) > 0 {
		return doc, nil
	}
	return nil, ErrNotFound
}
