package documentations

import (
	"net/http"

	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"github.com/TimoSto/ThesorTeX/services/website/internal/documentations"
)

func HandleDocumentations() func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		doc := r.URL.Query().Get("doc")
		lang := r.URL.Query().Get("lang")

		file, err := documentations.GetDoc(doc, lang)
		if err != nil {
			log.Error("could not read doc: %v", err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.Header().Set("Cache-Control", "max-age=3600")

		w.Write(file)
		//TODO: test
	}

	return fn
}
