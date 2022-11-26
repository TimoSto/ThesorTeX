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

func HandleAddProject(config conf.Config) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("createReqe")
		if r.Method != http.MethodPut {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Error("Error decoding data for adding project: %v", err) //todo: error-message-templates
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		project, err := projects.CreateProject(string(data), config, os.MkdirAll, ioutil.WriteFile)
		if err != nil {
			log.Error("Error creating project: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		pData, err := json.Marshal(project)
		if err != nil {
			log.Error("Error marshaling data: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(pData)
	}

	return http.HandlerFunc(fn)
}
