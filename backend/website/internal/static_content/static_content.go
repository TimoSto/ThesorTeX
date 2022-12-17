package static_content

import (
	"embed"
	"net/http"
)

//go:embed static
var static embed.FS

func Register(mux *http.ServeMux) {
	mux.Handle("/static/", http.FileServer(http.FS(static)))
}
