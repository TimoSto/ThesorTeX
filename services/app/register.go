package app

import (
	"fmt"
	"net/http"

	"github.com/TimoSto/ThesorTeX/services/app/conf"
	"github.com/TimoSto/ThesorTeX/services/app/internal/handlers"
)

const prefix = "/app"

func Register(mux *http.ServeMux, config conf.Config) {

	mux.Handle(fmt.Sprintf("%s/", prefix), handlers.HandleRoot())

	mux.Handle(fmt.Sprintf("%s/config", prefix), handlers.HandleConfig())

	mux.Handle(fmt.Sprintf("%s/projects", prefix), handlers.HandleProjects(config))

	mux.Handle(fmt.Sprintf("%s/createProject", prefix), handlers.HandleAddProject(config))

	mux.Handle(fmt.Sprintf("%s/deleteProject", prefix), handlers.HandleProjectDelete(config))

	mux.Handle(fmt.Sprintf("%s/projectData", prefix), handlers.HandleProjectData(config))

}
