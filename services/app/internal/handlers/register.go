package handlers

import (
	"net/http"

	"github.com/TimoSto/ThesorTeX/pkg/backend/roothandler"
	"github.com/TimoSto/ThesorTeX/services/app/internal/config"
)

func RegisterAppHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/", roothandler.GetRootHandler(config.Version))
}
