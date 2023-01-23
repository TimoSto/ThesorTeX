package entry

import (
	"net/http"

	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"github.com/TimoSto/ThesorTeX/services/app/internal/config"
	"github.com/TimoSto/ThesorTeX/services/app/internal/domain/entries"
	"github.com/TimoSto/ThesorTeX/services/app/internal/filesystem"
)

func HandleDeleteEntry(fs filesystem.FileSystem, cfg config.Config) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		query := r.URL.Query()

		if len(query["project"]) == 0 || len(query["key"]) == 0 {
			log.Error("missing query parameters")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err := entries.DeleteEntry(query["project"][0], query["key"][0], fs, cfg)
		if err != nil {
			log.Error("could not delete entry: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

	return fn
}
