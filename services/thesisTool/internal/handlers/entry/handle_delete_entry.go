package entry

import (
	"net/http"

	"github.com/TimoSto/ThesorTeX/pkg/backend/filesystem"
	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"github.com/TimoSto/ThesorTeX/services/thesisTool/internal/config"
	"github.com/TimoSto/ThesorTeX/services/thesisTool/internal/domain/entries"
)

func HandleDeleteEntry(fs filesystem.FileSystem) func(w http.ResponseWriter, r *http.Request) {
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

		err := entries.DeleteEntry(query["project"][0], query["key"][0], fs, config.Cfg)
		if err != nil {
			log.Error("could not delete entry: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

	return fn
}
