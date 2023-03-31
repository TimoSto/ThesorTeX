package documentations

import (
	"encoding/json"
	"net/http"

	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"github.com/TimoSto/ThesorTeX/services/website/internal/documentations_depr"
)

func HandleDocumentations() func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		doc := r.URL.Query().Get("doc")
		lang := r.URL.Query().Get("lang")

		var data []byte

		if doc == "thesis" {
			docObj, err := documentations.GetThesisDoc(lang)
			if err != nil {
				log.Error("could not read doc: %v", err)
				w.WriteHeader(http.StatusNotFound)
				return
			}
			data, err = json.Marshal(docObj)
			if err != nil {
				log.Error("could serialize doc: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		if doc == "thesis_tool" {
			docObj, err := documentations.GetThesisToolDoc(lang)
			if err != nil {
				log.Error("could not read doc: %v", err)
				w.WriteHeader(http.StatusNotFound)
				return
			}
			data, err = json.Marshal(docObj)
			if err != nil {
				log.Error("could serialize doc: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}

		w.Header().Set("Cache-Control", "max-age=3600")

		w.Write(data)
	}

	return fn
}
