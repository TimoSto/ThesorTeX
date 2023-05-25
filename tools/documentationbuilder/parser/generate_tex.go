package parser

import (
	"fmt"
	"strings"
)

func GenerateContentForTeX(title string, docs []DocBody) ([]byte, error) {
	body := fmt.Sprintf("\\part{%s}\n", title)

	for _, d := range docs {
		body += fmt.Sprintf("\\section{%s}\n", d.Title)
		for i, g := range d.Groups {
			if g.Type == "TEXT" {
				for _, e := range g.Elements {
					//e.Content = strings.Replace(e.Content, "\\anhang", "{{\\\\anhang}}", -1)
					//e.Content = strings.Replace(e.Content, "\\part", "{{\\\\part}}", -1)
					//e.Content = strings.Replace(e.Content, "\\section", "{{\\\\section}}", -1)
					e.Content = strings.Replace(e.Content, "\\", "\\textbackslash ", -1)
					e.Content = strings.Replace(e.Content, "_", "{{\\_}}", -1)
					body += fmt.Sprintf(getFormatForType(e.Style), e.Content)
				}
				if i < len(d.Groups)-1 && d.Groups[i+1].Type == "TEXT" {
					body += "\\\\"
				}
				body += "\n"
			} else if g.Type == "CODE" {
				body += "\\begin{verbatim}\n"
				for _, e := range g.Elements {
					body += e.Content
					body += "\n"
				}
				body += "\\end{verbatim}\n"
			} else if g.Type == "IMAGE" {
				//TODO: unit test
				body += "\\begin{figure}[H]\n"
				body += fmt.Sprintf("\\includegraphics[width=\\textwidth]{%s}\n", g.Elements[0].Content)
				body += fmt.Sprintf("\\caption{%s}\n", g.Elements[1].Content)
				body += "\\end{figure}\n"
				if i < len(d.Groups)-1 && d.Groups[i+1].Type == "TEXT" {
					body += "\\noindent "
				}
			}
		}
	}

	return []byte(body), nil
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
