package project

import (
	"encoding/json"
	"net/http"

	"github.com/TimoSto/ThesorTeX/backend/app/internal/database"
	"github.com/TimoSto/ThesorTeX/backend/pkg/log"
)

func HandleProjects(store database.ThesorTeXStore) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		projectsData, err := store.GetAllProjects()
		if err != nil {
			log.Error("Reading projects: %v", err)
		}

		data, err := json.Marshal(projectsData)
		if err != nil {
			log.Error("Marshaling projects: %v", err)
		}

		_, err = w.Write(data)
		if err != nil {
			log.Error("Sending projects data to client: %v", err)
			return
		}
	}

	return http.HandlerFunc(fn)
}
