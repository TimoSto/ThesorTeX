package category

import (
	"encoding/json"
	"net/http"

	"github.com/TimoSto/ThesorTeX/backend/pkg/log"
	bib_categories2 "github.com/TimoSto/ThesorTeX/services/app/internal/bib_categories"
	"github.com/TimoSto/ThesorTeX/services/app/internal/database"
)

type SaveData struct {
	Project        string
	Name           string
	InitialName    string
	CitaviCategory string
	CitaviFilter   []string
	BibFields      []bib_categories2.Field
	CiteFields     []bib_categories2.Field
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

		metaData, err := bib_categories2.SaveCategory(store, data.Project, data.Name, data.InitialName, data.CitaviCategory, data.CitaviFilter, data.BibFields, data.CiteFields)
		if err != nil {
			log.Error("got error saving the category: %v", err)
			w.WriteHeader(http.StatusBadRequest)
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
