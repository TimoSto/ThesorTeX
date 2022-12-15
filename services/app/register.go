package app

import (
	"fmt"
	"net/http"

	"github.com/TimoSto/ThesorTeX/services/app/conf"
	"github.com/TimoSto/ThesorTeX/services/app/internal/database/local_store"
	"github.com/TimoSto/ThesorTeX/services/app/internal/handlers/entry"
	"github.com/TimoSto/ThesorTeX/services/app/internal/handlers/project"
	"github.com/TimoSto/ThesorTeX/services/app/internal/handlers/root"
)

const prefix = "/app"

func Register(mux *http.ServeMux, config conf.Config) {

	store := local_store.Store{
		Config: config,
	}

	mux.Handle(fmt.Sprintf("%s/", prefix), root.HandleRoot())

	//mux.Handle(fmt.Sprintf("%s/config", prefix), handlers.HandleConfig())

	mux.Handle(fmt.Sprintf("%s/projects", prefix), project.HandleProjects(&store))

	mux.Handle(fmt.Sprintf("%s/createProject", prefix), project.HandleAddProject(&store))

	mux.Handle(fmt.Sprintf("%s/deleteProject", prefix), project.HandleProjectDelete(&store))

	mux.Handle(fmt.Sprintf("%s/projectData", prefix), project.HandleProjectData(&store))

	mux.Handle(fmt.Sprintf("%s/saveEntry", prefix), entry.HandleSaveEntry(config, &store))

}
