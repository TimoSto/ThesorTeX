package handlers

import (
	"net/http"

	"github.com/TimoSto/ThesorTeX/services/website/internal/handlers/assets"
	"github.com/TimoSto/ThesorTeX/services/website/internal/handlers/templates"
)

func RegisterWebsiteHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/", assets.HandleAssets())

	mux.HandleFunc("/templates/thesis", templates.HandleThesisTemplate())

	mux.HandleFunc("/templates/cv", templates.HandleCVTemplate())
}
