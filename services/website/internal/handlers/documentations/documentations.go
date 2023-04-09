package documentations

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"github.com/TimoSto/ThesorTeX/services/website/internal/documentations"
)

//TODO: unit test
func HandleDocumentations() func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		lang := r.URL.Query().Get("lang")
		format := r.URL.Query().Get("format")
		var data []byte
		var err error

		if format == "pdf" {
			data, err = documentations.GetPDFDocs(lang)
		} else {
			data, err = documentations.GetJsonDocs(lang)
		}

		if err != nil {
			if err == documentations.ErrNotFound {
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

func HandleDocumentationImages() func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		strippedPath := strings.Replace(r.URL.Path, "/documentation/images/", "", 1)
		fmt.Print(strippedPath)

		data, err := documentations.GetDocsImage(strippedPath)

		if err != nil {
			log.Error(err.Error())
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.Write(data)
	}

	return fn
}
