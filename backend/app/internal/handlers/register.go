package handlers

import (
	"fmt"
	"mime"
	"net/http"

	"github.com/TimoSto/ThesorTeX/backend/app/internal/conf"
	"github.com/TimoSto/ThesorTeX/backend/app/internal/database/local_store"
	"github.com/TimoSto/ThesorTeX/backend/app/internal/handlers/category"
	"github.com/TimoSto/ThesorTeX/backend/app/internal/handlers/config"
	"github.com/TimoSto/ThesorTeX/backend/app/internal/handlers/entry"
	"github.com/TimoSto/ThesorTeX/backend/app/internal/handlers/project"
	"github.com/TimoSto/ThesorTeX/backend/app/internal/handlers/root"
)

const prefix = "/app"

func Register(mux *http.ServeMux, cfg conf.Config) {

	store := local_store.Store{
		Config: cfg,
	}

	mime.AddExtensionType(".js", "text/javascript")
	mime.AddExtensionType(".css", "text/css")

	mux.Handle("/", root.HandleRoot())

	mux.Handle(fmt.Sprintf("%s/config", prefix), config.HandleConfig())

	mux.Handle(fmt.Sprintf("%s/projects", prefix), project.HandleProjects(&store))

	mux.Handle(fmt.Sprintf("%s/createProject", prefix), project.HandleAddProject(&store))

	mux.Handle(fmt.Sprintf("%s/deleteProject", prefix), project.HandleProjectDelete(&store))

	mux.Handle(fmt.Sprintf("%s/projectData", prefix), project.HandleProjectData(&store))

	mux.Handle(fmt.Sprintf("%s/renameProject", prefix), project.HandleSaveProjectMetaData(&store))

	mux.Handle(fmt.Sprintf("%s/saveEntry", prefix), entry.HandleSaveEntry(&store))

	mux.Handle(fmt.Sprintf("%s/deleteEntry", prefix), entry.HandleDeleteEntry(&store))

	mux.Handle(fmt.Sprintf("%s/saveCategory", prefix), category.HandleSaveCategory(&store))

	mux.Handle(fmt.Sprintf("%s/deleteCategory", prefix), category.HandleDeleteCategory(&store))

}
