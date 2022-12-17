package entry

import (
	"encoding/json"
	"net/http"

	"github.com/TimoSto/ThesorTeX/backend/app/internal/bib_entries"
	"github.com/TimoSto/ThesorTeX/backend/app/internal/database"
	"github.com/TimoSto/ThesorTeX/pkg/log"
)

type SaveEntryData struct {
	Project    string
	InitialKey string
	Key        string
	Category   string
	Fields     []string
}

func HandleSaveEntry(store database.ThesorTeXStore) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var data SaveEntryData
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&data)
		if err != nil {
			log.Error("got error reading data of entry save: %v", err)
			w.WriteHeader(http.StatusBadRequest)
		}

		entries := []database.BibEntry{
			{
				Key:        data.Key,
				Category:   data.Category,
				Fields:     data.Fields,
				Comment:    "",
				CiteNumber: 0,
			},
		}

		err = bib_entries.SaveEntriesToProject(data.Project, store, entries, []string{data.InitialKey})

		if err != nil {
			log.Error("Error saving entry: %v", err)
		}
	}

	return http.HandlerFunc(fn)
}
