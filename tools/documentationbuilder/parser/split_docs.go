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
var mainTitleRegex = regexp.MustCompile("^MainTitle: .*\n")

func SplitDocs(file string) (string, []RawDocs) {
	var docs []RawDocs

	mainTitleMatch := mainTitleRegex.FindString(file)
	mainTitle := strings.Trim(strings.TrimLeft(mainTitleMatch, "MainTitle: "), "\n")

	if mainTitle == "" {
		panic("no title was found")
	}

	fmt.Println(mainTitleMatch)

	file = file[len(mainTitleMatch):]
	fmt.Println(file)

	splitted := splitRegex.Split(file, -1)

	for _, s := range splitted {
		s = strings.Trim(s, "\n")
		titleMatch := titleRegex.FindString(s)

		docs = append(docs, RawDocs{
			Title:   strings.Trim(strings.TrimLeft(titleMatch, "title: "), "\n"),
			Content: strings.Trim(strings.TrimLeft(s, strings.Trim(titleMatch, "\n")), "\n"),
		})
	}

	return mainTitle, docs
}
