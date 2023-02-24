package examples

import (
	"net/http"

	"github.com/TimoSto/ThesorTeX/services/website/internal/example_project"
)

func HandleExamples() func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Write(example_project.ExampleDE)
	}

	return fn
}
