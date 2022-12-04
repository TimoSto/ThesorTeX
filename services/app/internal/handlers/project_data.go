package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/TimoSto/ThesorTeX/pkg/log"
	"github.com/TimoSto/ThesorTeX/services/app/conf"
	"github.com/TimoSto/ThesorTeX/services/app/internal/bib_categories"
	"github.com/TimoSto/ThesorTeX/services/app/internal/bib_entries"
)

type ProjectData struct {
	Entries    []bib_entries.BibEntry
	Categories []bib_categories.BibCategory
}

func HandleProjectData(config conf.Config) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		project := r.FormValue("project")
		if len(project) == 0 {
			log.Error("missing url query param 'project'")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		entries, err := bib_entries.ReadEntriesOfProject(project, ioutil.ReadFile, config)
		if err != nil {
			log.Error(fmt.Sprintf("failed to read entries for project %s: %v", project, err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		categories, err := bib_categories.ReadCategoriesOfProject(project, ioutil.ReadFile, config)
		if err != nil {
			log.Error(fmt.Sprintf("failed to read categories for project %s: %v", project, err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data, err := json.Marshal(ProjectData{
			Entries:    entries,
			Categories: categories,
		})
		if err != nil {
			log.Error(fmt.Sprintf("failed to marshal data for project %s: %v", project, err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = w.Write(data)
		if err != nil {
			log.Error(fmt.Sprintf("failed to send data for project %s: %v", project, err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	return http.HandlerFunc(fn)
}
