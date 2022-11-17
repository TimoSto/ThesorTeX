package register

import (
	"net/http"

	"github.com/TimoSto/ThesorTeX/services/project_management/internal/handlers"
)

func Register(mux *http.ServeMux) {

	mux.Handle("/", handlers.HandleRoot())

}
