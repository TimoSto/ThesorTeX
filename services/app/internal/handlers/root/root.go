package root

import (
	"fmt"
	"net/http"

	"github.com/TimoSto/ThesorTeX/services/app/frontend/assets"
	"github.com/TimoSto/ThesorTeX/services/app/internal/conf"
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
