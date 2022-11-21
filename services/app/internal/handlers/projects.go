package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/TimoSto/ThesorTeX/pkg/log"
	"github.com/TimoSto/ThesorTeX/services/app/conf"
	"github.com/TimoSto/ThesorTeX/services/app/internal/projects"
)

func HandleProjects(config conf.Config) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		projectsData, err := projects.GetAllProjects(config, ioutil.ReadDir, os.Mkdir, ioutil.ReadFile)
		if err != nil {
			log.Error(fmt.Sprintf("Reading projects: %v", err))
		}

		data, err := json.Marshal(projectsData)
		if err != nil {
			log.Error(fmt.Sprintf("Marshaling projects: %v", err))
		}

		w.Write(data)
	}

	return http.HandlerFunc(fn)
}
