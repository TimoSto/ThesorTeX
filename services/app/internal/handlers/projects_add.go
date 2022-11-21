package handlers

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/TimoSto/ThesorTeX/pkg/log"
	"github.com/TimoSto/ThesorTeX/services/app/conf"
	"github.com/TimoSto/ThesorTeX/services/app/internal/projects"
)

func HandleAddProject(config conf.Config) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Error("Error decoding data for adding project: %v", err) //todo: error-message-templates
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = projects.CreateProject(string(data), config, os.MkdirAll, ioutil.WriteFile)
		if err != nil {
			log.Error("Error creating project: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	return http.HandlerFunc(fn)
}
