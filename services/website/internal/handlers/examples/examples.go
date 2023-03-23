package examples

import (
	"net/http"

	"github.com/TimoSto/ThesorTeX/services/website/internal/example_project"
)

func HandleExamples() func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		lang := r.URL.Query().Get("lang")
		format := r.URL.Query().Get("format")

		var example []byte
		if format == "pdf" {
			example = example_project.GetExamplePDF(lang)
		} else {
			example = example_project.GetExampleZip(lang)
		}

		w.Write(example)
	}

	return fn
}
