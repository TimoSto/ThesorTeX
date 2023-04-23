package faviconhandler

import (
	_ "embed"
	"fmt"
	"net/http"
	"time"
)

//TODO: generate this at build time
//go:embed favicon.svg
var favicon []byte

func GetFaviconHandler() func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", fmt.Sprintf("max-age=%v", int(24*time.Hour/time.Second)))

		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		w.Header().Set("Content-Type", "image/svg+xml")

		w.Write(favicon)
	}

	return fn
}
