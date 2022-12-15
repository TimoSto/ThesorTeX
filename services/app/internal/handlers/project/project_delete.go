package project

import (
	"io/ioutil"
	"net/http"

	"github.com/TimoSto/ThesorTeX/pkg/log"
	"github.com/TimoSto/ThesorTeX/services/app/internal/database"
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

		err = store.DeleteProject(string(data))

		if err != nil {
			log.Error("Error deleting project: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	return http.HandlerFunc(fn)
}
