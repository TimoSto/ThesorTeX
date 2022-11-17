package assets

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed dist
var dist embed.FS

var assets, _ = fs.Sub(dist, "dist")

type AssetConf struct {
	Ignore    []string
	MapRootTo string
}

func Register(mux *http.ServeMux, c AssetConf) {

	fn := func(w http.ResponseWriter, r *http.Request) {

		for _, p := range c.Ignore {
			if r.URL.Path == p {
				w.WriteHeader(http.StatusNotFound)
				return
			}
		}

		if r.URL.Path == "/" {
			r.URL.Path = c.MapRootTo
		}

		http.FileServer(http.FS(assets)).ServeHTTP(w, r)
	}

	mux.Handle("/", http.HandlerFunc(fn))
}
