package documentations

import (
	_ "embed"
	"errors"
	"fmt"
	"strings"
)

//go:embed docs/thesisTemplate/de/parsed.json
var templateDocJSON_DE []byte

var (
	ErrNotFound = errors.New("documentation not found")
)

var docsMap = map[string][]byte{
	"thesisTemplate_JSON_DE": templateDocJSON_DE,
}

func GetDocumentation(topic string, format string, lang string) ([]byte, error) {
	doc := docsMap[fmt.Sprintf("%s_%s_%s", topic, format, strings.ToUpper(lang))]

	if len(doc) > 0 {
		return doc, nil
	}
	return nil, ErrNotFound
}
