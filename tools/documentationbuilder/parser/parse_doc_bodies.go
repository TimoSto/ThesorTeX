package parser

import (
	"fmt"
	"strings"
)

type DocBody struct {
	Groups []group
}

type group struct {
	Type     string
	Elements []element
}

type element struct {
	Content string
	Style   string
}

func ParseDocBodies(raw []RawDocs) []DocBody {
	var bodies []DocBody

	for _, r := range raw {
		bodies = append(bodies, parseDocBody(r))
	}

	return bodies
}

func parseDocBody(raw RawDocs) DocBody {
	var body DocBody

	body.Groups = []group{
		{},
	}

	splitted := strings.Split(raw.Content, "\n")

	for _, s := range splitted {
		l := analyseLine(s)
		fmt.Println(l)
	}

	return body
}

type AllowedType string

const (
	TypeText  AllowedType = "Text"
	TypeEmpty             = "Empty"
)

type analyseLineResult struct {
	Type    AllowedType
	Content string
}

func analyseLine(line string) analyseLineResult {
	if line == "" || line == "\n" {
		return analyseLineResult{
			Type:    TypeEmpty,
			Content: "",
		}
	}

	return analyseLineResult{
		Type:    TypeText,
		Content: line,
	}
}
