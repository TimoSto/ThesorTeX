package documentations

import (
	"embed"
)

//go:embed docs
var docs embed.FS

type Doc struct {
	Title   string
	Content string
}

// TODO: add "title: ..." to all md files and then display that as title of the expandable area
type Docs struct {
	Docs []Doc
}

var thesisPaths = []string{
	"thesis_template_usage",
	"thesis_template_usage/chapter_numbering",
	"thesis_template_usage/header_footer",
	"thesis_template_usage/abbreviations",
	"thesis_template_usage/appendix",
	"thesis_template_usage/bibliography",
}

func GetThesisDoc(lang string) (Docs, error) {
	var doc Docs

	for _, p := range thesisPaths {
		d, err := readDoc(p, lang)
		if err != nil {
			return doc, err
		}
		doc.Docs = append(doc.Docs, Doc{
			Title:   d.Title,
			Content: d.Content,
		})
	}

	return doc, nil
}

var thesisToolPaths = []string{
	"thesis_tool_usage",
	"thesis_tool_usage/startup",
}

func GetThesisToolDoc(lang string) (Docs, error) {
	var doc Docs

	for _, p := range thesisToolPaths {
		d, err := readDoc(p, lang)
		if err != nil {
			return doc, err
		}
		doc.Docs = append(doc.Docs, Doc{
			Title:   d.Title,
			Content: d.Content,
		})
	}

	return doc, nil
}
