package assets

import (
	"net/http"

	"github.com/TimoSto/ThesorTeX/services/website/frontend/assets"
)

func HandleAssets() func(w http.ResponseWriter, r *http.Request) {
	httpFs := http.FileServer(http.FS(assets.Assets))
	fn := func(w http.ResponseWriter, r *http.Request) {
		httpFs.ServeHTTP(w, r)
	}

	return fn
}
