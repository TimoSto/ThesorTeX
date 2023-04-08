package builder

import (
	"fmt"
	goFs "io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/TimoSto/ThesorTeX/pkg/backend/project_template"
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

//TODO: use file system
func BuildDocumentationFromTemplate(outPath string, body string) error {
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

			content = r.ReplaceAllString(content, body)
			b = []byte(content)
		}

		fmt.Println("Writing: ", filepath.Join(outPath, path))

		return os.WriteFile(filepath.Join(outPath, path), b, os.ModePerm)
	})
}
