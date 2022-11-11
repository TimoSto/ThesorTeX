package static_content

import (
	"embed"
	"net/http"

	"github.com/TimoSto/ThesorTeX/pkg/server"
)

//go:embed static
var static embed.FS

func Register(srv server.Server) {
	srv.Register("/static/", func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.FS(static)).ServeHTTP(w, r)
	})
}
