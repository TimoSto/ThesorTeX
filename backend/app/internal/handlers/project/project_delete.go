package project

import (
	"io/ioutil"
	"net/http"

	"github.com/TimoSto/ThesorTeX/backend/app/internal/database"
	"github.com/TimoSto/ThesorTeX/backend/pkg/log"
)

func HandleProjectDelete(store database.ThesorTeXStore) http.Handler {
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

		err = store.RemoveProject(string(data))

		if err != nil {
			log.Error("Error deleting project: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	return http.HandlerFunc(fn)
}
