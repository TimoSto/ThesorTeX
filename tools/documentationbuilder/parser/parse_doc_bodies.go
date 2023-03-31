package parser

import (
	"fmt"
	"regexp"
	"strings"
)

type DocBody struct {
	Groups []group
}

type group struct {
	Type     string
	Elements []element
}

const (
	StylePlain         = "Plain"
	StyleBold          = "Bold"
	StyleItalic        = "Italic"
	StyleItalicAndBold = "ItalicAndBold"
)

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

//TODO: these regex also match a leading space for some reason
var (
	boldAndItalicRegex = regexp.MustCompile("[^*]\\*{3}[^***]+\\*{3}")
	boldRegex          = regexp.MustCompile("[^*]\\*{2}[^**]+\\*{2}")
	italicRegex        = regexp.MustCompile("[^*]\\*[^*]+\\*")
)

func splitLineIntoElements(line string) []element {
	// find italic and bold parts
	var elements []element

	line = " " + line

	matches := italicRegex.FindAllString(line, -1)
	matches = append(matches, boldRegex.FindAllString(line, -1)...)
	matches = append(matches, boldAndItalicRegex.FindAllString(line, -1)...)

	matchIndexes := italicRegex.FindAllStringIndex(line, -1)
	matchIndexes = append(matchIndexes, boldRegex.FindAllStringIndex(line, -1)...)
	matchIndexes = append(matchIndexes, boldAndItalicRegex.FindAllStringIndex(line, -1)...)

	matchIndexes, matches = sortMatches(matchIndexes, matches)

	if len(matchIndexes) == 0 {
		elements = []element{
			{
				Style:   StylePlain,
				Content: line[1:],
			},
		}
	}

	beg := 0
	end := 0
	for i, match := range matchIndexes {

		// add content before match as plain text
		end = match[0]
		if i == 0 {
			if string(line[beg]) == " " {
				beg++
			}
		}
		if string(line[end]) == " " {
			end++
		}

		plainContentToAdd := line[beg:end]
		if plainContentToAdd != "" {
			elements = append(elements, element{
				Content: plainContentToAdd,
				Style:   StylePlain,
			})
		}

		// add the actual result
		matchValue := line[match[0] : match[0]+len(matches[i])]
		shift := 0
		if string(matchValue[0]) == " " {
			shift++
		}

		if italicRegex.MatchString(matchValue) {
			elements = append(elements, element{
				Content: matchValue[1+shift : len(matchValue)-1],
				Style:   StyleItalic,
			})
		}
		if boldRegex.MatchString(matchValue) {
			elements = append(elements, element{
				Content: matchValue[2+shift : len(matchValue)-2],
				Style:   StyleBold,
			})
		}
		if boldAndItalicRegex.MatchString(matchValue) {
			elements = append(elements, element{
				Content: matchValue[3+shift : len(matchValue)-3],
				Style:   StyleItalicAndBold,
			})
		}

		beg = end + len(matchValue) - 1

		//end = match[0]
		//if match[1] != 0 {
		//	elements = append(strings, s[beg:end])
		//}
		//beg = match[1]
		//// This also appends the current match
		//strings = append(strings, s[match[0]:match[1]])
	}

	if len(matchIndexes) > 0 && beg < len(line) {
		elements = append(elements, element{
			Content: line[beg:],
			Style:   StylePlain,
		})
	}

	return elements
}

func sortMatches(indexes [][]int, values []string) ([][]int, []string) {
	// bubble sort
	for i := 0; i < len(indexes)-1; i++ {
		for j := 0; j < len(indexes)-i-1; j++ {
			if indexes[j][0] > indexes[j+1][0] {
				indexes[j], indexes[j+1] = indexes[j+1], indexes[j]
				values[j], values[j+1] = values[j+1], values[j]
			}
		}
	}
	return indexes, values
}
