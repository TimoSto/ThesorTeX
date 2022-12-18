package project

import (
	"encoding/json"
	"net/http"

	"github.com/TimoSto/ThesorTeX/backend/app/internal/database"
	"github.com/TimoSto/ThesorTeX/backend/app/internal/project"
	"github.com/TimoSto/ThesorTeX/pkg/log"
)

type renameData struct {
	Project string
	Data    database.ProjectMetaData
}

func HandleSaveProjectMetaData(store database.ThesorTeXStore) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var data renameData
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&data)
		if err != nil {
			log.Error("got error reading data of project rename: %v", err)
			w.WriteHeader(http.StatusBadRequest)
		}

		err = project.SaveProjectMetaData(data.Project, data.Data, store)
		if err != nil {
			log.Error("got error saving project meta data: %v", err)
			w.WriteHeader(http.StatusBadRequest)
		}
	}

	return http.HandlerFunc(fn)
}
