package documentations

import (
	"embed"
	"fmt"
)

//go:embed docs
var docs embed.FS

type ThesisDoc struct {
	Main             string
	ChapterNumbering string
	HeaderFooter     string
	Abbreviations    string
	Appendix         string
}

func GetThesisDoc(lang string) (ThesisDoc, error) {
	var doc ThesisDoc

	val, err := docs.ReadFile(fmt.Sprintf("docs/%s/%s.md", "thesis_template_usage", lang))
	if err != nil {
		return doc, err
	}

	doc.Main = string(val)

	val, err = docs.ReadFile(fmt.Sprintf("docs/%s/%s.md", "thesis_template_usage/chapter_numbering", lang))
	if err != nil {
		return doc, err
	}

	doc.ChapterNumbering = string(val)

	val, err = docs.ReadFile(fmt.Sprintf("docs/%s/%s.md", "thesis_template_usage/header_footer", lang))
	if err != nil {
		return doc, err
	}

	doc.HeaderFooter = string(val)

	val, err = docs.ReadFile(fmt.Sprintf("docs/%s/%s.md", "thesis_template_usage/abbreviations", lang))
	if err != nil {
		return doc, err
	}

	doc.Abbreviations = string(val)

	val, err = docs.ReadFile(fmt.Sprintf("docs/%s/%s.md", "thesis_template_usage/appendix", lang))
	if err != nil {
		return doc, err
	}

	doc.Appendix = string(val)

	return doc, nil
}
