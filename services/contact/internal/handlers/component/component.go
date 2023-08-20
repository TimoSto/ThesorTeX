package component

import (
	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"github.com/TimoSto/ThesorTeX/services/contact/frontend/assets"
	"io"
	"net/http"
	"path/filepath"
)

func GetComponentHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		path := r.URL.Path[len("/contact/feedbackComponent/"):]

		file, err := assets.Assets.Open(path)
		if err != nil {
			log.Error("could not read component file: %v", err)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		b, err := io.ReadAll(file)
		if err != nil {
			log.Error("could not read component file: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")

		ext := filepath.Ext(path)
		switch ext {
		case ".js":
			w.Header().Set("Content-Type", "text/javascript")
			break
		case ".css":
			w.Header().Set("Content-Type", "text/css")
			break
		}

		_, err = w.Write(b)
		if err != nil {
			log.Error("could not send component file: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

	return http.HandlerFunc(fn)
}
