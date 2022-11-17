package assets

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed dist
var dist embed.FS

var assets, _ = fs.Sub(dist, "dist")

func Register(mux *http.ServeMux) {

	fn := func(w http.ResponseWriter, r *http.Request) {

		http.FileServer(http.FS(assets)).ServeHTTP(w, r)
	}

	mux.Handle("/", http.HandlerFunc(fn))
}
