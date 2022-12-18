package category

import (
	"encoding/json"
	"net/http"

	"github.com/TimoSto/ThesorTeX/backend/app/internal/bib_categories"
	"github.com/TimoSto/ThesorTeX/backend/app/internal/database"
	"github.com/TimoSto/ThesorTeX/pkg/log"
)

type deleteCategoryData struct {
	Project string
	Name    string
}

func HandleDeleteCategory(store database.ThesorTeXStore) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var data deleteCategoryData
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&data)
		if err != nil {
			log.Error("got error reading data of category delete: %v", err)
			w.WriteHeader(http.StatusBadRequest)
		}

		err = bib_categories.DeleteCategory(data.Project, data.Name, store)
		if err != nil {
			log.Error("got error deleting the categroy: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

	return http.HandlerFunc(fn)
}
