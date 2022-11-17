package app

import (
	"fmt"
	"net/http"

	"github.com/TimoSto/ThesorTeX/services/app/internal/handlers"
)

const prefix = "/app"

func Register(mux *http.ServeMux) {

	mux.Handle(fmt.Sprintf("%s/", prefix), handlers.HandleRoot())

}
