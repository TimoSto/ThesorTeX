package handlers

import (
	"net/http"

	"github.com/TimoSto/ThesorTeX/pkg/backend/versionhandler"
	"github.com/TimoSto/ThesorTeX/services/app/internal/config"
	"github.com/TimoSto/ThesorTeX/services/app/internal/filesystem"
	"github.com/TimoSto/ThesorTeX/services/app/internal/handlers/assets"
	"github.com/TimoSto/ThesorTeX/services/app/internal/handlers/category"
	"github.com/TimoSto/ThesorTeX/services/app/internal/handlers/entry"
	"github.com/TimoSto/ThesorTeX/services/app/internal/handlers/project"
)

func RegisterAppHandlers(mux *http.ServeMux, fs filesystem.FileSystem, cfg config.Config) {
	mux.HandleFunc("/", assets.HandleAssets())

	mux.HandleFunc("/version", versionhandler.GetRootHandler(config.Version))

	mux.HandleFunc("/createNewProject", project.CreateProjectHandler(fs, cfg))

	mux.HandleFunc("/getAllProjects", project.GetAllProjectsHandler(fs, cfg))

	mux.HandleFunc("/getProjectData", project.GetProjectDataHandler(fs, cfg))

	mux.HandleFunc("/deleteProject", project.HandleProjectDelete(fs, cfg))

	mux.HandleFunc("/saveCategory", category.HandleSaveCategory(fs, cfg))

	mux.HandleFunc("/deleteCategory", category.HandleDeleteCategory(fs, cfg))

	mux.HandleFunc("/saveEntry", entry.HandleSaveEntry(fs, cfg))

	mux.HandleFunc("/uploadEntries", entry.HandleUploadEntries(fs, cfg))

	mux.HandleFunc("/deleteEntry", entry.HandleDeleteEntry(fs, cfg))
}
