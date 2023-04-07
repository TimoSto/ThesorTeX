package parser

import (
	"fmt"
	goFs "io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/TimoSto/ThesorTeX/pkg/backend/project_template"
)

//TODO: use file system
func BuildDocumentationFromTemplate(outPath string, docs []DocBody) error {
	outPath = filepath.Join(outPath, "tex")

	err := os.MkdirAll(outPath, os.ModePerm)
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Join(outPath, "data"), os.ModePerm)
	if err != nil {
		return err
	}

	err = os.MkdirAll(filepath.Join(outPath, "styPackages"), os.ModePerm)
	if err != nil {
		return err
	}

	return goFs.WalkDir(project_template.ProjectTemplate, ".", func(path string, d goFs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		b, err := goFs.ReadFile(project_template.ProjectTemplate, path)
		if err != nil {
			return err // or panic or ignore
		}

		path = strings.TrimPrefix(path, "template/")

		//TODO: unit test
		if path == "main.tex" {
			content := string(b)
			r := regexp.MustCompile("(?s)% Start Content(.*?)% End Content")

			c, err := GenerateContentForTeX(docs)
			if err != nil {
				panic(err)
			}
			content = r.ReplaceAllString(content, string(c))
			b = []byte(content)
		}

		fmt.Println("Writing: ", filepath.Join(outPath, path))

		return os.WriteFile(filepath.Join(outPath, path), b, os.ModePerm)
	})
}

func GenerateContentForTeX(docs []DocBody) ([]byte, error) {
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
