package handlers

import (
	"net/http"

	"github.com/TimoSto/ThesorTeX/pkg/backend/roothandler"
	"github.com/TimoSto/ThesorTeX/services/app/internal/config"
	"github.com/TimoSto/ThesorTeX/services/app/internal/filesystem"
	"github.com/TimoSto/ThesorTeX/services/app/internal/handlers/category"
	"github.com/TimoSto/ThesorTeX/services/app/internal/handlers/project"
)

func RegisterAppHandlers(mux *http.ServeMux, fs filesystem.FileSystem, cfg config.Config) {
	mux.HandleFunc("/", roothandler.GetRootHandler(config.Version))

	mux.HandleFunc("/createNewProject", project.CreateProjectHandler(fs, cfg))

	mux.HandleFunc("/getAllProjects", project.GetAllProjectsHandler(fs, cfg))

	mux.HandleFunc("/getProjectData", project.GetProjectDataHandler(fs, cfg))

	mux.HandleFunc("/deleteProject", project.HandleProjectDelete(fs, cfg))

	mux.HandleFunc("/saveCategory", category.HandleSaveCategory(fs, cfg))
}
