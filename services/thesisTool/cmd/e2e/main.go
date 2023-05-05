package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/TimoSto/ThesorTeX/pkg/backend/filesystem/local"
	"github.com/TimoSto/ThesorTeX/pkg/backend/handler_chain"
	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"github.com/TimoSto/ThesorTeX/pkg/backend/pathbuilder"
	"github.com/TimoSto/ThesorTeX/services/thesisTool/internal/config"
	"github.com/TimoSto/ThesorTeX/services/thesisTool/internal/domain/projects"
	"github.com/TimoSto/ThesorTeX/services/thesisTool/internal/domain/updatechecker"
	"github.com/TimoSto/ThesorTeX/services/thesisTool/internal/handlers"
)

func main() {
	loglevel := os.Getenv("LOG_LEVEL")
	if loglevel == "" {
		loglevel = "INFO"
	}
	log.Setup(false, loglevel)

	config.Version = "0.0.1"
	log.Info("Version: %s", config.Version)

	log.Info("Starting the e2e app...")

	//pathbuilder.Init()

	mux := http.NewServeMux()

	chain := handler_chain.CreateHandlerChain()

	fs := local.FileSystem{}

	_, err := config.ReadConfig()
	if err != nil {
		log.Error("could read/init config: %v", err)
		os.Exit(1)
	}

	log.Setup(false, config.Cfg.LogLevel)

	updatechecker.CheckUpdateAvailable()

	envPort := os.Getenv("E2E_PORT")

	if envPort != "" {
		config.Cfg.Port = envPort
	} else {
		config.Cfg.Port = "8440"
	}

	config.Cfg.ProjectsDir = "projects"

	log.Info("removing possibly existing previous projects...")
	err = os.RemoveAll(pathbuilder.GetPathFromExecRoot("projects"))
	if err != nil {
		log.Error("could not delete projects dir: %v", err)
		os.Exit(1)
	}

	log.Info("Creating example project...")
	_, err = projects.CreateProject("example", &fs, config.Cfg)
	if err != nil {
		log.Error("unexpected error creating the example project: %v", err)
		os.Exit(1)
	}
	log.Info("Created example project under %s", pathbuilder.GetProjectPath(config.Cfg.ProjectsDir, "example"))

	handlers.RegisterAppHandlers(mux, &fs)

	err = http.ListenAndServe(fmt.Sprintf("localhost:%s", config.Cfg.Port), chain.Then(mux))
	if err != nil {
		log.Error("unexpected error starting server: %v", err)
		os.Exit(1)
	}
}
