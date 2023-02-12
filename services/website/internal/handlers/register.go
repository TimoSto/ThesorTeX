package handlers

import (
	"net/http"

	"github.com/TimoSto/ThesorTeX/services/website/internal/handlers/assets"
	"github.com/TimoSto/ThesorTeX/services/website/internal/handlers/templates"
	"github.com/TimoSto/ThesorTeX/services/website/internal/handlers/versions"
)

func RegisterWebsiteHandlers(mux *http.ServeMux, dev bool) {
	mux.HandleFunc("/", assets.HandleAssets())

	mux.HandleFunc("/templates/thesis", templates.HandleThesisTemplate())

	mux.HandleFunc("/templates/cv", templates.HandleCVTemplate())

	mux.HandleFunc("/versions/tool", versions.HandleToolVersions(dev))
}
