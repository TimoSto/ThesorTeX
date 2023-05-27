package category

import (
	"encoding/json"
	"github.com/TimoSto/ThesorTeX/services/thesisTool/internal/domain/entries"
	"net/http"

	"github.com/TimoSto/ThesorTeX/pkg/backend/filesystem"
	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"github.com/TimoSto/ThesorTeX/services/thesisTool/internal/config"
	"github.com/TimoSto/ThesorTeX/services/thesisTool/internal/domain/categories"
)

func HandleSaveCategory(fs filesystem.FileSystem) func(w http.ResponseWriter, r *http.Request) {
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

		project := query["project"][0]
		ogName := query["name"][0]
		err = categories.SaveCategory(fs, config.Cfg, project, ogName, data)
		if err != nil {
			log.Error("could not save category: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = entries.UpdateCategoryReferences(fs, config.Cfg, project, ogName, data.Name)
		if err != nil {
			log.Error("could not update category references in entries: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	return fn
}
