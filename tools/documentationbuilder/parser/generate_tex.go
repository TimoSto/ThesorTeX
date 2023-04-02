package parser

import "fmt"

func GenerateContentForTeX(docs []DocBody) (string, error) {
	var body string

	for _, d := range docs {
		body += fmt.Sprintf("\\section{%s}\n", d.Title)
		for i, g := range d.Groups {
			if g.Type == "TEXT" {
				for _, e := range g.Elements {
					body += fmt.Sprintf(getFormatForType(e.Style), e.Content)
				}
				if i < len(d.Groups)-1 {
					body += "\\\\"
				}
				body += "\n"
			}
		}
	}

	return body, nil
}

func getFormatForType(t string) string {
	if t == "BOLD" {
		return "\\textbf{%s}"
	}
	if t == "ITALIC" {
		return "\\textit{%s}"
	}
	if t == "ITALIC-BOLD" {
		return "\\textit{\\textbf{%s}}"
	}

	return "%s"
}
