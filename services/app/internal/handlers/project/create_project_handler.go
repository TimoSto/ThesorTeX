package project

import (
	"encoding/json"
	"net/http"

	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"github.com/TimoSto/ThesorTeX/services/app/internal/config"
	"github.com/TimoSto/ThesorTeX/services/app/internal/domain/projects"
	"github.com/TimoSto/ThesorTeX/services/app/internal/filesystem"
)

type createProjectData struct {
	Name string
}

func CreateProjectHandler(fs filesystem.FileSystem) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var data createProjectData
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&data)
		if err != nil {
			log.Error("got error reading data of project creation: %v", err)
			w.WriteHeader(http.StatusBadRequest)
		}

		metaData, err := projects.CreateProject(data.Name, fs, config.Cfg)

		if err != nil {
			log.Error("got error creating new project: %v", err)
			w.WriteHeader(http.StatusBadRequest)
		}

		respData, err := json.Marshal(metaData)
		if err != nil {
			log.Error("got error sending data of created project back to client: %v", err)
			w.WriteHeader(http.StatusBadRequest)
		}

		w.Write(respData)
	}

	return fn
}
