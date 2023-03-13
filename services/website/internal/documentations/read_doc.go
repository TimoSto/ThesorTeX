package documentations

import (
	"fmt"
	"regexp"
	"strings"
)

func readDoc(path string, lang string) (Doc, error) {
	var doc Doc

	val, err := docs.ReadFile(fmt.Sprintf("docs/%s/%s.md", path, lang))
	if err != nil {
		return doc, err
	}

	titleRegex := regexp.MustCompile("^title: .*\n")
	titleMatch := titleRegex.FindString(string(val))

	doc.Title = strings.Trim(strings.TrimLeft(titleMatch, "title: "), "\n")
	doc.Content = strings.Trim(strings.TrimLeft(string(val), titleMatch), "\n")

	return doc, nil
}
