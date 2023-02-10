package templates

import (
	"net/http"

	"github.com/TimoSto/ThesorTeX/services/website/internal/templates"
)

func HandleCVTemplate() func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		tmpl := templates.CVTemplate

		w.Write(tmpl)
	}

	return fn
}
