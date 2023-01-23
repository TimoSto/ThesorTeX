package entry

import (
	"net/http"

	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"github.com/TimoSto/ThesorTeX/services/app/internal/config"
	"github.com/TimoSto/ThesorTeX/services/app/internal/filesystem"
)

func HandleSaveEntry(fs filesystem.FileSystem, cfg config.Config) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		query := r.URL.Query()

		if len(query["project"]) == 0 || len(query["key"]) == 0 {
			log.Error("missing query parameters")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	return fn
}
