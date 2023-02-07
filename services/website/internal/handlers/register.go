package handlers

import (
	"net/http"

	"github.com/TimoSto/ThesorTeX/services/website/internal/handlers/assets"
)

func RegisterWebsiteHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/", assets.HandleAssets())
}
