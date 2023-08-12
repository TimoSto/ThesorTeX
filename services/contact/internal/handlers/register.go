package handlers

import (
	"github.com/TimoSto/ThesorTeX/pkg/backend/versionhandler"
	"github.com/TimoSto/ThesorTeX/services/contact/internal/config"
	"github.com/TimoSto/ThesorTeX/services/contact/internal/feedback"
	"github.com/TimoSto/ThesorTeX/services/contact/internal/handlers/sendfeedback"
	"net/http"
)

func RegisterHandlers(mux *http.ServeMux, store feedback.Store) {
	mux.HandleFunc("/contact/version", versionhandler.GetRootHandler(config.Version))

	mux.Handle("/contact/feedback", sendfeedback.GetFeedbackHandler(store))
}
