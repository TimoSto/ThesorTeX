package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/TimoSto/ThesorTeX/pkg/log"
	"github.com/TimoSto/ThesorTeX/services/app/internal/database"
	"github.com/TimoSto/ThesorTeX/services/app/internal/project"
)

func HandleProjectData(store database.ThesorTeXStore) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		name := r.FormValue("project")
		if len(name) == 0 {
			log.Error("missing url query param 'name'")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		entries, err := store.GetProjectEntries(name)
		if err != nil {
			log.Error("failed to read entries for project %s: %v", name, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		categories, err := store.GetProjectCategories(name)
		if err != nil {
			log.Error("failed to read categories for project %s: %v", name, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		projectData := project.ProjectData{
			Entries:    entries,
			Categories: categories,
		}

		data, err := json.Marshal(projectData)
		if err != nil {
			log.Error("failed to marshal data for name %s: %v", name, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = w.Write(data)
		if err != nil {
			log.Error("failed to send data for name %s: %v", name, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	return http.HandlerFunc(fn)
}
