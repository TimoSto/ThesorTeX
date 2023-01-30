package category

import (
	"net/http"

	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"github.com/TimoSto/ThesorTeX/services/app/internal/config"
	"github.com/TimoSto/ThesorTeX/services/app/internal/domain/categories"
	"github.com/TimoSto/ThesorTeX/services/app/internal/filesystem"
)

func HandleDeleteCategory(fs filesystem.FileSystem) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		query := r.URL.Query()

		if len(query["project"]) == 0 || len(query["name"]) == 0 {
			log.Error("missing query parameters")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err := categories.DeleteCategory(query["project"][0], query["name"][0], fs, config.Cfg)
		if err != nil {
			log.Error("could not delete category: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

	return fn
}
