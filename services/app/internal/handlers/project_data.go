package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/TimoSto/ThesorTeX/pkg/log"
	"github.com/TimoSto/ThesorTeX/services/app/internal/database"
)

type ProjectData struct {
}

func HandleProjectData(store database.ThesorTeXStore) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		project := r.FormValue("project")
		if len(project) == 0 {
			log.Error("missing url query param 'project'")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		projectData, err := store.GetProjectData(project)
		if err != nil {
			log.Error("failed to read data for project %s: %v", project, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data, err := json.Marshal(projectData)
		if err != nil {
			log.Error("failed to marshal data for project %s: %v", project, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = w.Write(data)
		if err != nil {
			log.Error("failed to send data for project %s: %v", project, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	return http.HandlerFunc(fn)
}
