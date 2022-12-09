package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/TimoSto/ThesorTeX/pkg/log"
	"github.com/TimoSto/ThesorTeX/services/app/conf"
	"github.com/TimoSto/ThesorTeX/services/app/internal/bib_entries"
)

type SaveEntryData struct {
	Project    string
	InitialKey string
	Key        string
	Category   string
	Fields     []string
}

func HandleSaveEntry(config conf.Config) http.Handler {
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

		entries := []bib_entries.BibEntry{
			{
				Key:        data.Key,
				Category:   data.Category,
				Fields:     data.Fields,
				Comment:    "",
				CiteNumber: 0,
			},
		}

		err = bib_entries.SaveEntriesToProject(data.Project, ioutil.ReadFile, ioutil.WriteFile, config, entries, []string{data.InitialKey})

		if err != nil {
			log.Error("Error saving entry: %v", err)
		}
	}

	return http.HandlerFunc(fn)
}
