package category

import (
	"encoding/json"
	"net/http"

	"github.com/TimoSto/ThesorTeX/backend/app/internal/bib_categories"
	"github.com/TimoSto/ThesorTeX/backend/app/internal/database"
	"github.com/TimoSto/ThesorTeX/pkg/log"
)

type SaveData struct {
	Project        string
	Name           string
	InitialName    string
	CitaviCategory string
	CitaviFilter   []string
	BibFields      []database.Field
	CiteFields     []database.Field
}

func HandleSaveCategory(store database.ThesorTeXStore) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var data SaveData
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&data)
		if err != nil {
			log.Error("got error reading data of category save: %v", err)
			w.WriteHeader(http.StatusBadRequest)
		}

		err = bib_categories.SaveCategory(store, data.Project, data.Name, data.InitialName, data.CitaviCategory, data.CitaviFilter, data.BibFields, data.CiteFields)
		if err != nil {
			log.Error("got error saving the category: %v", err)
			w.WriteHeader(http.StatusBadRequest)
		}
	}

	return http.HandlerFunc(fn)
}
