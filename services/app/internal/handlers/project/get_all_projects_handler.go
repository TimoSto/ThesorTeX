package project

import (
	"encoding/json"
	"net/http"

	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"github.com/TimoSto/ThesorTeX/services/app/internal/config"
	"github.com/TimoSto/ThesorTeX/services/app/internal/domain/projects"
	"github.com/TimoSto/ThesorTeX/services/app/internal/filesystem"
)

func GetAllProjectsHandler(fs filesystem.FileSystem, cfg config.Config) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		p, err := projects.GetAllProjects(fs, cfg)
		if err != nil {
			log.Error("got unexpected error reading all projects: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data, err := json.Marshal(p)
		if err != nil {
			log.Error("got unexpected error formatting data for all projects: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(data)

	}

	return fn
}
