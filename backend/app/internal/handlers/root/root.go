package root

import (
	"fmt"
	"net/http"

	"github.com/TimoSto/ThesorTeX/backend/app/internal/conf"
	"github.com/TimoSto/ThesorTeX/frontend/assets"
)

func HandleRoot() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			w.Write([]byte(fmt.Sprintf(`{"Version": "%s"}`, conf.Version)))
			return
		}
		http.FileServer(http.FS(assets.SubDir)).ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
