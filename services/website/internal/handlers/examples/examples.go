package examples

import (
	"net/http"

	"github.com/TimoSto/ThesorTeX/services/website/internal/example_project"
)

func HandleExamples() func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		lang := r.URL.Query().Get("lang")

		example := example_project.GetExampleZip(lang)

		w.Write(example)
	}

	return fn
}
