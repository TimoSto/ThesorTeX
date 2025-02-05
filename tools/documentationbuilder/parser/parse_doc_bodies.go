package parser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type DocBody struct {
	Title     string
	Groups    []group
	Footnotes map[int][]element
}

type AllowedType string

const (
	TypeText        AllowedType = "Text"
	TypeEmpty                   = "Empty"
	TypeBeginCode               = "BeginCode"
	TypeCode                    = "Code"
	TypeEndCode                 = "EndCode"
	TypeListItem                = "ListItem"
	TypeImage                   = "Image"
	TypeImageLabel              = "ImageLabel"
	TypeFootnoteRef             = "FootnoteRef"
	TypeSubSection              = "Subsection"
)

type group struct {
	Type     AllowedType
	Elements []element
}

const (
	StylePlain         = "PLAIN"
	StyleBold          = "BOLD"
	StyleItalic        = "ITALIC"
	StyleItalicAndBold = "ITALIC-BOLD"
	LinkTitle          = "LINK-TITLE"
	LinkHref           = "LINK-HREF"
	Footnote           = "FOOTNOTE"
)

type element struct {
	Content string
	Style   string
}

func ParseDocBodies(raw []RawDocs) []DocBody {
	var bodies []DocBody

	for _, r := range raw {
		if r.Content != "" {
			bodies = append(bodies, parseDocBody(r))
		}
	}

	return bodies
}

var imgLabelRegex = regexp.MustCompile("\\[(.*?)\\]")
var imgLinkRegex = regexp.MustCompile("\\((.*?)\\)")

func parseDocBody(raw RawDocs) DocBody {
	body := DocBody{
		Title:  raw.Title,
		Groups: []group{{}},
	}

	splitted := strings.Split(raw.Content, "\n")

	lengthElements := 1

	incode := false

	for _, s := range splitted {
		l := analyseLine(s, incode)
		fmt.Println(l.Type)
		if l.Type == TypeEmpty && !incode {
			// empty line => if the last group in the docs already has elements, start a new group
			if len(body.Groups[len(body.Groups)-1].Elements) > 0 {
				body.Groups = append(body.Groups, group{})
				lengthElements = 1
			}
		} else if l.Type == TypeText {
			if body.Groups[len(body.Groups)-1].Type == "" {
				//first element in group sets the type
				body.Groups[len(body.Groups)-1].Type = "TEXT"
			} else if body.Groups[len(body.Groups)-1].Type != "TEXT" {
				// if type differs with previous, create a new group
				body.Groups = append(body.Groups, group{Type: "TEXT"})
				lengthElements = 1
			} else {
				// still text but new line => space
				body.Groups[len(body.Groups)-1].Elements[lengthElements-1].Content += " "
			}

			elements := splitLineIntoElements(l.Content)

			sortElementsIntoGroups(&body, elements, &lengthElements)
		} else if l.Type == TypeSubSection {
			body.Groups = append(body.Groups, group{Type: "SUBSECTION", Elements: []element{
				{
					Style:   "",
					Content: l.Content,
				},
			}})
			lengthElements = 1
		} else if l.Type == TypeBeginCode {
			incode = true
			body.Groups = append(body.Groups, group{Type: "CODE"})
			lengthElements = 1
		} else if l.Type == TypeEndCode {
			incode = false
		} else if l.Type == TypeCode {
			body.Groups[len(body.Groups)-1].Elements = append(body.Groups[len(body.Groups)-1].Elements, element{
				Content: l.Content,
			})
		} else if l.Type == TypeListItem {
			if body.Groups[len(body.Groups)-1].Type != "LIST" {
				// if type differs with previous, create a new group
				body.Groups = append(body.Groups, group{Type: "LIST"})
				lengthElements = 1
			}

			//TODO: support bold, italic, ... link in regular text
			body.Groups[len(body.Groups)-1].Elements = append(body.Groups[len(body.Groups)-1].Elements, element{
				Content: l.Content,
			})
		} else if l.Type == TypeImage {
			label := imgLabelRegex.FindString(l.Content)
			link := imgLinkRegex.FindString(l.Content)
			body.Groups = append(body.Groups, group{Type: "IMAGE"})
			lengthElements = 1
			// image group always has these two elements
			body.Groups[len(body.Groups)-1].Elements = []element{
				{
					Content: link[1 : len(link)-1],
					Style:   "",
				},
				{
					Content: label[1 : len(label)-1],
					Style:   "",
				},
			}
		} else if l.Type == TypeFootnoteRef {

			num := footnoteRefRegex.FindString(l.Content)

			content := l.Content[len(num):]

			num = num[2 : len(num)-2]

			// TODO: handle error
			n, _ := strconv.Atoi(num)

			if body.Footnotes == nil {
				body.Footnotes = make(map[int][]element)
			}

			elements := splitLineIntoElements(content)

			body.Footnotes[n] = elements
		}
	}

	if len(body.Groups[len(body.Groups)-1].Elements) == 0 {
		// a newline before footnotes causes an empty group
		// TODO: do this better
		body.Groups = body.Groups[:len(body.Groups)-1]
	}

	return body
}

type analyseLineResult struct {
	Type    AllowedType
	Content string
}

var imgRegex = regexp.MustCompile("!\\[[^\\]]*\\]\\([^\\)]*\\)")
var footnoteRefRegex = regexp.MustCompile("^\\[\\^[0-9]{1,3}\\]:")
var subSectionRegex = regexp.MustCompile("^### .*")

func analyseLine(line string, incode bool) analyseLineResult {
	if line == "" || line == "\n" {
		return analyseLineResult{
			Type:    TypeEmpty,
			Content: "",
		}
	}
	if strings.Trim(line, " ") == "```latex" {
		return analyseLineResult{
			Type:    TypeBeginCode,
			Content: "",
		}
	}
	if strings.Trim(line, " ") == "```" {
		return analyseLineResult{
			Type:    TypeEndCode,
			Content: "",
		}
	}
	if len(line) > 2 && line[:2] == "- " {
		return analyseLineResult{
			Type:    TypeListItem,
			Content: line[2:],
		}
	}

	img := imgRegex.FindString(line)
	if img != "" {
		return analyseLineResult{
			Type:    TypeImage,
			Content: img,
		}
	}

	if incode {
		return analyseLineResult{
			Type:    TypeCode,
			Content: line,
		}
	}

	if footnoteRefRegex.MatchString(line) {
		return analyseLineResult{
			Type:    TypeFootnoteRef,
			Content: line,
		}
	}

	if subSectionRegex.MatchString(line) {
		return analyseLineResult{
			Type:    TypeSubSection,
			Content: strings.TrimLeft(line, "### "),
		}
	}

	return analyseLineResult{
		Type:    TypeText,
		Content: line,
	}
}

// TODO: these regex also match a leading space for some reason
var (
	boldAndItalicRegex = regexp.MustCompile("[^*]\\*{3}[^***]+\\*{3}")
	boldRegex          = regexp.MustCompile("[^*]\\*{2}[^**]+\\*{2}")
	italicRegex        = regexp.MustCompile("[^*]\\*[^*]+\\*")
	linkRegex          = regexp.MustCompile("\\[[^\\]]*\\]\\([^)]*\\)*")
	footnoteRegex      = regexp.MustCompile("\\[\\^[0-9]{1,3}\\]")
)

func splitLineIntoElements(line string) []element {
	// find italic and bold parts
	var elements []element

	line = " " + line

	matches := italicRegex.FindAllString(line, -1)
	matches = append(matches, boldRegex.FindAllString(line, -1)...)
	matches = append(matches, boldAndItalicRegex.FindAllString(line, -1)...)
	matches = append(matches, linkRegex.FindAllString(line, -1)...)
	matches = append(matches, footnoteRegex.FindAllString(line, -1)...)

	matchIndexes := italicRegex.FindAllStringIndex(line, -1)
	matchIndexes = append(matchIndexes, boldRegex.FindAllStringIndex(line, -1)...)
	matchIndexes = append(matchIndexes, boldAndItalicRegex.FindAllStringIndex(line, -1)...)
	matchIndexes = append(matchIndexes, linkRegex.FindAllStringIndex(line, -1)...)
	matchIndexes = append(matchIndexes, footnoteRegex.FindAllStringIndex(line, -1)...)

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
		if string(line[end]) != "*" && string(line[end]) != "[" {
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
		if string(matchValue[0]) != "*" {
			shift++
		}

		if italicRegex.MatchString(matchValue) {
			elements = append(elements, element{
				Content: matchValue[1+shift : len(matchValue)-1],
				Style:   StyleItalic,
			})
		} else if boldRegex.MatchString(matchValue) {
			elements = append(elements, element{
				Content: matchValue[2+shift : len(matchValue)-2],
				Style:   StyleBold,
			})
		} else if boldAndItalicRegex.MatchString(matchValue) {
			elements = append(elements, element{
				Content: matchValue[3+shift : len(matchValue)-3],
				Style:   StyleItalicAndBold,
			})
		} else if linkRegex.MatchString(matchValue) {
			title := matchValue[1:strings.Index(matchValue, "]")]
			href := matchValue[strings.Index(matchValue, "(")+1 : strings.Index(matchValue, ")")]
			if title == "" {
				title = href
			}
			elements = append(elements, element{
				Content: title,
				Style:   LinkTitle,
			})
			elements = append(elements, element{
				Content: href,
				Style:   LinkHref,
			})
			//TODO: why is this necessary?
			end++
		} else if footnoteRegex.MatchString(matchValue) {
			val := strings.TrimLeft(matchValue, "[^")
			val = strings.TrimRight(val, "]")

			elements = append(elements, element{
				Content: val,
				Style:   Footnote,
			})
			//TODO: why is this necessary? Do I need to change the regex?
			end++
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

func sortElementsIntoGroups(body *DocBody, elements []element, lengthElements *int) {
	for _, e := range elements {
		if len(body.Groups[len(body.Groups)-1].Elements) < *lengthElements {
			// if tto few elements are present, create an empty one
			body.Groups[len(body.Groups)-1].Elements = append(
				body.Groups[len(body.Groups)-1].Elements, element{})
		}
		if body.Groups[len(body.Groups)-1].Elements[*lengthElements-1].Style == "" {
			// first element in group sets style
			body.Groups[len(body.Groups)-1].Elements[*lengthElements-1].Style = e.Style
			body.Groups[len(body.Groups)-1].Elements[*lengthElements-1].Content = e.Content
		} else if body.Groups[len(body.Groups)-1].Elements[*lengthElements-1].Style == e.Style {
			// if style matches, jsut add the content
			body.Groups[len(body.Groups)-1].Elements[*lengthElements-1].Content += e.Content
		} else {
			//if style differs, create new group and add in there
			body.Groups[len(body.Groups)-1].Elements = append(
				body.Groups[len(body.Groups)-1].Elements, element{})
			*lengthElements++
			body.Groups[len(body.Groups)-1].Elements[*lengthElements-1].Style = e.Style
			body.Groups[len(body.Groups)-1].Elements[*lengthElements-1].Content = e.Content
		}
	}
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
