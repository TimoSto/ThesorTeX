package entry

import (
	"encoding/json"
	"net/http"

	"github.com/TimoSto/ThesorTeX/backend/app/internal/bib_entries"
	"github.com/TimoSto/ThesorTeX/backend/app/internal/database"
	"github.com/TimoSto/ThesorTeX/backend/pkg/log"
)

type deleteEntryData struct {
	Project string
	Key     string
}

func HandleDeleteEntry(store database.ThesorTeXStore) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var data deleteEntryData
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&data)
		if err != nil {
			log.Error("got error reading data of entry delete: %v", err)
			w.WriteHeader(http.StatusBadRequest)
		}

		err = bib_entries.DeleteEntry(data.Project, data.Key, store)
		if err != nil {
			log.Error("got error deleting the entry: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

	return http.HandlerFunc(fn)
}
