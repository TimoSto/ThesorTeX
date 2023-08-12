package handlers

import (
	"fmt"
	"github.com/TimoSto/ThesorTeX/pkg/backend/versionhandler"
	"github.com/TimoSto/ThesorTeX/services/contact/internal/config"
	"github.com/TimoSto/ThesorTeX/services/contact/internal/feedback"
	"github.com/TimoSto/ThesorTeX/services/contact/internal/handlers/sendfeedback"
	"net/http"
)

const prefix = "/contact"

func RegisterHandlers(mux *http.ServeMux, store feedback.Store) {
	mux.HandleFunc(fmt.Sprintf("%s/version", prefix), versionhandler.GetRootHandler(config.Version))

	mux.Handle(fmt.Sprintf("%s/contact/feedback", prefix), sendfeedback.GetFeedbackHandler(store))
}
