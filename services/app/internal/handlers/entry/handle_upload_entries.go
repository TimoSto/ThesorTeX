package entry

import (
	"encoding/json"
	"net/http"

	"github.com/TimoSto/ThesorTeX/pkg/backend/filesystem"
	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"github.com/TimoSto/ThesorTeX/services/app/internal/config"
	"github.com/TimoSto/ThesorTeX/services/app/internal/domain/entries"
)

func HandleUploadEntries(fs filesystem.FileSystem) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		query := r.URL.Query()

		if len(query["project"]) == 0 {
			log.Error("missing query parameter")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var uploadedEntries []entries.Entry
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&uploadedEntries)
		if err != nil {
			log.Error("cannot decode entries to save: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = entries.SaveEntries(fs, config.Cfg, query["project"][0], uploadedEntries)
		if err != nil {
			log.Error("got error saving uploaded entries: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

	return fn
}
