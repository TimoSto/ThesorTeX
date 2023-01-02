package handlers

import (
	"fmt"
	"mime"
	"net/http"

	"github.com/TimoSto/ThesorTeX/services/app/internal/conf"
	"github.com/TimoSto/ThesorTeX/services/app/internal/database/local_store"
	category2 "github.com/TimoSto/ThesorTeX/services/app/internal/handlers/category"
	"github.com/TimoSto/ThesorTeX/services/app/internal/handlers/config"
	entry2 "github.com/TimoSto/ThesorTeX/services/app/internal/handlers/entry"
	project2 "github.com/TimoSto/ThesorTeX/services/app/internal/handlers/project"
	"github.com/TimoSto/ThesorTeX/services/app/internal/handlers/root"
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

	mux.Handle(fmt.Sprintf("%s/projects", prefix), project2.HandleProjects(&store))

	mux.Handle(fmt.Sprintf("%s/createProject", prefix), project2.HandleAddProject(&store))

	mux.Handle(fmt.Sprintf("%s/deleteProject", prefix), project2.HandleProjectDelete(&store))

	mux.Handle(fmt.Sprintf("%s/projectData", prefix), project2.HandleProjectData(&store))

	mux.Handle(fmt.Sprintf("%s/renameProject", prefix), project2.HandleSaveProjectMetaData(&store))

	mux.Handle(fmt.Sprintf("%s/saveEntry", prefix), entry2.HandleSaveEntry(&store))

	mux.Handle(fmt.Sprintf("%s/deleteEntry", prefix), entry2.HandleDeleteEntry(&store))

	mux.Handle(fmt.Sprintf("%s/saveCategory", prefix), category2.HandleSaveCategory(&store))

	mux.Handle(fmt.Sprintf("%s/deleteCategory", prefix), category2.HandleDeleteCategory(&store))

}
