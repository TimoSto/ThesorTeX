package category

import (
	"encoding/json"
	"net/http"

	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"github.com/TimoSto/ThesorTeX/services/app/internal/config"
	"github.com/TimoSto/ThesorTeX/services/app/internal/domain/categories"
	"github.com/TimoSto/ThesorTeX/services/app/internal/filesystem"
)

func HandleSaveCategory(fs filesystem.FileSystem, cfg config.Config) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		query := r.URL.Query()

		if len(query["project"]) == 0 || len(query["name"]) == 0 {
			log.Error("missing query parameters")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var data categories.Category
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&data)
		if err != nil {
			log.Error("cannot decode category to save: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = categories.SaveCategory(fs, cfg, query["project"][0], query["name"][0], data)
		if err != nil {
			log.Error("could not save category: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	return fn
}