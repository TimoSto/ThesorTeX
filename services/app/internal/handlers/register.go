package handlers

import (
	"net/http"

	"github.com/TimoSto/ThesorTeX/pkg/backend/faviconhandler"
	"github.com/TimoSto/ThesorTeX/pkg/backend/versionhandler"
	"github.com/TimoSto/ThesorTeX/services/app/internal/config"
	"github.com/TimoSto/ThesorTeX/services/app/internal/filesystem"
	"github.com/TimoSto/ThesorTeX/services/app/internal/handlers/assets"
	"github.com/TimoSto/ThesorTeX/services/app/internal/handlers/category"
	config_handlers "github.com/TimoSto/ThesorTeX/services/app/internal/handlers/config"
	"github.com/TimoSto/ThesorTeX/services/app/internal/handlers/entry"
	"github.com/TimoSto/ThesorTeX/services/app/internal/handlers/project"
)

func RegisterAppHandlers(mux *http.ServeMux, fs filesystem.FileSystem) {
	mux.HandleFunc("/", assets.HandleAssets())

	mux.HandleFunc("/version", versionhandler.GetRootHandler(config.Version))

	mux.HandleFunc("/favicon", faviconhandler.GetFaviconHandler())

	mux.HandleFunc("/createNewProject", project.CreateProjectHandler(fs))

	mux.HandleFunc("/getAllProjects", project.GetAllProjectsHandler(fs))

	mux.HandleFunc("/getProjectData", project.GetProjectDataHandler(fs))

	mux.HandleFunc("/deleteProject", project.HandleProjectDelete(fs))

	mux.HandleFunc("/saveCategory", category.HandleSaveCategory(fs))

	mux.HandleFunc("/deleteCategory", category.HandleDeleteCategory(fs))

	mux.HandleFunc("/saveEntry", entry.HandleSaveEntry(fs))

	mux.HandleFunc("/uploadEntries", entry.HandleUploadEntries(fs))

	mux.HandleFunc("/deleteEntry", entry.HandleDeleteEntry(fs))

	mux.HandleFunc("/getConfig", config_handlers.HandleConfigGet())

	mux.HandleFunc("/saveConfig", config_handlers.HandleConfigSave())
}
