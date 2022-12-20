package root

import (
	"net/http"

	"github.com/TimoSto/ThesorTeX/frontend/assets"
)

func HandleRoot() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.FS(assets.SubDir)).ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
