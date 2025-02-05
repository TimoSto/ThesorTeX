package project

import (
	"encoding/json"
	"net/http"

	"github.com/TimoSto/ThesorTeX/pkg/backend/filesystem"
	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"github.com/TimoSto/ThesorTeX/services/thesisTool/internal/config"
	"github.com/TimoSto/ThesorTeX/services/thesisTool/internal/domain/categories"
	"github.com/TimoSto/ThesorTeX/services/thesisTool/internal/domain/entries"
)

type ProjectData struct {
	Entries    []entries.Entry
	Categories []categories.Category
}

func GetProjectDataHandler(fs filesystem.FileSystem) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		query := r.URL.Query()

		if len(query["project"]) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		project := query["project"][0]

		projectEntries, err := entries.GetAllEntries(project, fs, config.Cfg)
		if err != nil {
			log.Error("got error reading entries of project from request: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		projectCategories, err := categories.GetAllCategories(project, fs, config.Cfg)
		if err != nil {
			log.Error("got error reading entries of project from request: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		obj := ProjectData{
			Entries:    projectEntries,
			Categories: projectCategories,
		}

		respData, err := json.Marshal(obj)
		if err != nil {
			log.Error("got error sending data back: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.Write(respData)
	}

	return fn
}
