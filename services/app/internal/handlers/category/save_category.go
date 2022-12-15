package category

import (
	"net/http"

	"github.com/TimoSto/ThesorTeX/services/app/internal/database"
)

func HandleSaveCategory(store database.ThesorTeXStore) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {}

	return http.HandlerFunc(fn)
}
