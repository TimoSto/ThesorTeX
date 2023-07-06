package handlers

import (
	"github.com/TimoSto/ThesorTeX/pkg/backend/versionhandler"
	"github.com/TimoSto/ThesorTeX/services/contact/internal/config"
	"github.com/TimoSto/ThesorTeX/services/contact/internal/feedback"
	"net/http"
)

func RegisterHandlers(mux *http.ServeMux, store feedback.Store) {
	mux.HandleFunc("/version", versionhandler.GetRootHandler(config.Version))
}
