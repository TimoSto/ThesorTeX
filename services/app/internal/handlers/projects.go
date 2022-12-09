package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/TimoSto/ThesorTeX/pkg/log"
	"github.com/TimoSto/ThesorTeX/services/app/conf"
	"github.com/TimoSto/ThesorTeX/services/app/internal/projects"
)

func HandleProjects(config conf.Config) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		projectsData, err := projects.GetAllProjects(config, ioutil.ReadDir, os.Mkdir, ioutil.ReadFile)
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
