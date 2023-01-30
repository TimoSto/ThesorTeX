package project

import (
	"net/http"

	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"github.com/TimoSto/ThesorTeX/services/app/internal/config"
	"github.com/TimoSto/ThesorTeX/services/app/internal/domain/projects"
	"github.com/TimoSto/ThesorTeX/services/app/internal/filesystem"
)

func HandleProjectDelete(fs filesystem.FileSystem) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		query := r.URL.Query()
		if len(query["project"]) == 0 {
			log.Error("Missing project query parameter")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err := projects.DeleteProject(query["project"][0], fs, config.Cfg)
		if err != nil {
			log.Error("unexpected error, deleting project: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	return fn
}
