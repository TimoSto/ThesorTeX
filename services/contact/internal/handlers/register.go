package handlers

import (
	"github.com/TimoSto/ThesorTeX/pkg/backend/versionhandler"
	"github.com/TimoSto/ThesorTeX/services/contact/internal/config"
	"net/http"
)

func RegisterHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/version", versionhandler.GetRootHandler(config.Version))
}
