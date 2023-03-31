package documentations

import (
	"net/http"

	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	documentations2 "github.com/TimoSto/ThesorTeX/services/website/internal/documentations"
)

func HandleDocumentations() func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		doc := r.URL.Query().Get("doc")
		lang := r.URL.Query().Get("lang")
		format := r.URL.Query().Get("format")

		data, err := documentations2.GetDocumentation(doc, format, lang)

		if err != nil {
			if err == documentations2.ErrNotFound {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			log.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Cache-Control", "max-age=3600")

		w.Write(data)
	}

	return fn
}
