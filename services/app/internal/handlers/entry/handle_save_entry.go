package entry

import (
	"encoding/json"
	"net/http"

	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"github.com/TimoSto/ThesorTeX/services/app/internal/config"
	"github.com/TimoSto/ThesorTeX/services/app/internal/domain/entries"
	"github.com/TimoSto/ThesorTeX/services/app/internal/filesystem"
)

func HandleSaveEntry(fs filesystem.FileSystem) func(w http.ResponseWriter, r *http.Request) {
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

		var data entries.Entry
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&data)
		if err != nil {
			log.Error("cannot decode entry to save: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = entries.SaveEntry(fs, config.Cfg, query["project"][0], query["key"][0], data)
		if err != nil {
			log.Error("could not save entry: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	return fn
}
