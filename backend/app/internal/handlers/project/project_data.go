package project

import (
	"encoding/json"
	"net/http"

	"github.com/TimoSto/ThesorTeX/backend/app/internal/bib_categories"
	"github.com/TimoSto/ThesorTeX/backend/app/internal/bib_entries"
	"github.com/TimoSto/ThesorTeX/backend/app/internal/database"
	"github.com/TimoSto/ThesorTeX/backend/pkg/log"
)

type ProjectData struct {
	Entries    []bib_entries.BibEntry
	Categories []bib_categories.BibCategory
}

func HandleProjectData(store database.ThesorTeXStore) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		name := r.FormValue("project")
		if len(name) == 0 {
			log.Error("missing url query param 'name'")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		entries, err := bib_entries.GetSortedEntries(name, store)
		if err != nil {
			log.Error("failed to read entries for project %s: %v", name, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		categories, err := bib_categories.GetSortedCategories(name, store)
		if err != nil {
			log.Error("failed to read categories for project %s: %v", name, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		projectData := ProjectData{
			Entries:    entries,
			Categories: categories,
		}

		data, err := json.Marshal(projectData)
		if err != nil {
			log.Error("failed to marshal data for name %s: %v", name, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = w.Write(data)
		if err != nil {
			log.Error("failed to send data for name %s: %v", name, err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	return http.HandlerFunc(fn)
}
