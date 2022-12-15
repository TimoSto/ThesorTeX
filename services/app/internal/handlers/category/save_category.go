package category

import (
	"encoding/json"
	"net/http"

	"github.com/TimoSto/ThesorTeX/pkg/log"
	"github.com/TimoSto/ThesorTeX/services/app/internal/database"
)

type SaveData struct {
	Project        string
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
			log.Error("got error reading data of entry save: %v", err)
			w.WriteHeader(http.StatusBadRequest)
		}
	}

	return http.HandlerFunc(fn)
}
