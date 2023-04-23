package faviconhandler

import (
	_ "embed"
	"net/http"
)

//TODO: generate this at build time
//go:embed favicon.svg
var favicon []byte

func GetFaviconHandler() func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "image/svg+xml")

		w.Write(favicon)
	}

	return fn
}
