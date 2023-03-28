package parser

import (
	"fmt"
	"regexp"
	"strings"
)

type RawDocs struct {
	Title   string
	Content string
}

var splitRegex = regexp.MustCompile("(?m)^---$")
var titleRegex = regexp.MustCompile("^title: .*\n")

func SplitDocs(file string) []RawDocs {
	var docs []RawDocs

	splitted := splitRegex.Split(file, -1)

	for _, s := range splitted {
		s = strings.Trim(s, "\n")
		fmt.Println("splitted", s)
		titleMatch := titleRegex.FindString(s)
		fmt.Println("title", titleMatch)

		docs = append(docs, RawDocs{
			Title:   strings.Trim(strings.TrimLeft(titleMatch, "title: "), "\n"),
			Content: strings.Trim(strings.TrimLeft(s, strings.Trim(titleMatch, "\n")), "\n"),
		})
	}

	return docs
}
