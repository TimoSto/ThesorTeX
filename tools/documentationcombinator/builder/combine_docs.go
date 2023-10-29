package builder

import (
	"fmt"
	goFs "io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/TimoSto/ThesorTeX/pkg/tex/project_template"
)

func CombineDocs(paths []string) (string, error) {
	var docs string

	//TODO use fake fs and test
	for _, p := range paths {
		f, err := os.ReadFile(filepath.Join(p, "parsed.tex"))
		if err != nil {
			return docs, err
		}
		docs += fmt.Sprintf("\n%s", string(f))
	}

	return docs, nil
}

// TODO: use file system
func BuildDocumentationFromTemplate(outPath string, body string, titlepage string, title string, author string, lang string, stripParts []string) error {
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

			for _, p := range stripParts {
				switch p {
				case "abbreviations":
					rabbr := regexp.MustCompile("(?s)% Change the title of the list of abbreviations here(.*?)\\\\clearpage\n\n")
					content = rabbr.ReplaceAllString(content, "")
					break
				case "listoffigures":
					rabbr := regexp.MustCompile("(?s)% Change the title of the list of figures here(.*?)\\\\clearpage\n\n")
					content = rabbr.ReplaceAllString(content, "")
					break
				case "listoftables":
					rabbr := regexp.MustCompile("(?s)% Change the title of the list of tables here(.*?)\\\\clearpage\n\n")
					content = rabbr.ReplaceAllString(content, "")
					break
				case "appendix":
					rabbr := regexp.MustCompile("(?s)% start appendix(.*?)% end appendix\n")
					content = rabbr.ReplaceAllString(content, "")
					break
				case "bibliography":
					rabbr := regexp.MustCompile("(?s)% Start bibliography(.*?)% end bibliography\n")
					content = rabbr.ReplaceAllString(content, "")
					break
				}
			}

			if lang != "DE" {
				content = strings.Replace(content, "\\renewcommand{\\figurename}{Abb.}", "\\renewcommand{\\figurename}{Fig.}", 1)
			}
			if titlepage != "" {
				content = strings.Replace(content, "%\\includepdf[pages={1-}]{titlepage.pdf}", fmt.Sprintf("\\includepdf[pages={1-}]{%s}", titlepage), 1)
			}
			if title != "" {
				content = strings.Replace(content, "\\renewcommand{\\mytitle}{The title of my document\\\\possibly with multiple lines}", fmt.Sprintf("\\renewcommand{\\mytitle}{%s}", title), 1)
			}
			if author != "" {
				content = strings.Replace(content, "\\renewcommand{\\myauthor}{My name}", fmt.Sprintf("\\renewcommand{\\myauthor}{%s}", author), 1)
			}

			r := regexp.MustCompile("(?s)% Start Content(.*?)% End Content")

			content = r.ReplaceAllString(content, body)
			b = []byte(content)
		}

		fmt.Println("Writing: ", filepath.Join(outPath, path))

		return os.WriteFile(filepath.Join(outPath, path), b, os.ModePerm)
	})
}
