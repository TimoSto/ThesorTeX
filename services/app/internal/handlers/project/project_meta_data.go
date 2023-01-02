package project

import (
	"encoding/json"
	"net/http"

	"github.com/TimoSto/ThesorTeX/pkg/log"
	"github.com/TimoSto/ThesorTeX/services/app/internal/database"
	project2 "github.com/TimoSto/ThesorTeX/services/app/internal/project"
)

type renameData struct {
	Project string
	Data    project2.ProjectMetaData
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

		err = project2.SaveProjectMetaData(data.Project, data.Data, store)
		if err != nil {
			log.Error("got error saving project meta data: %v", err)
			w.WriteHeader(http.StatusBadRequest)
		}
	}

	return http.HandlerFunc(fn)
}
