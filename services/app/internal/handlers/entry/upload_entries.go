package entry

import (
	"encoding/json"
	"net/http"

	"github.com/TimoSto/ThesorTeX/pkg/log"
	"github.com/TimoSto/ThesorTeX/services/app/internal/bib_entries"
	"github.com/TimoSto/ThesorTeX/services/app/internal/database"
)

type uploadData struct {
	Project string
	Entries []bib_entries.BibEntry
}

func HandleUploadEntries(store database.ThesorTeXStore) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var data uploadData
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&data)
		if err != nil {
			log.Error("got error reading data of entry save: %v", err)
			w.WriteHeader(http.StatusBadRequest)
		}

		metaData, err := bib_entries.ApplyEntries(data.Project, store, data.Entries, nil)

		if err != nil {
			log.Error("Error saving entries: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		respData, err := json.Marshal(metaData)
		if err != nil {
			log.Error("Error marshaling project meta data: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.Write(respData)
	}

	return http.HandlerFunc(fn)
}
