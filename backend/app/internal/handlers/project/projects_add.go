package project

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/TimoSto/ThesorTeX/backend/app/internal/database"
	"github.com/TimoSto/ThesorTeX/backend/app/internal/project"
	"github.com/TimoSto/ThesorTeX/pkg/log"
)

func HandleAddProject(store database.ThesorTeXStore) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
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

		project, err := project.CreateProject(string(data), store)
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
