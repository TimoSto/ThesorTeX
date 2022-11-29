package handlers

import (
	"io/ioutil"
	"net/http"

	"github.com/TimoSto/ThesorTeX/pkg/log"
	"github.com/TimoSto/ThesorTeX/services/app/conf"
	"github.com/TimoSto/ThesorTeX/services/app/internal/projects"
)

func HandleProjectDelete(config conf.Config) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Error("Error decoding data for deleting project: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = projects.DeleteProject(string(data), config)

		if err != nil {
			log.Error("Error deleting project: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	return http.HandlerFunc(fn)
}
