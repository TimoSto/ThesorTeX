package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TimoSto/ThesorTeX/pkg/log"
	"github.com/TimoSto/ThesorTeX/services/app/internal/projects"
)

func HandleProjects() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		projectsData, err := projects.GetAllProjects()
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
