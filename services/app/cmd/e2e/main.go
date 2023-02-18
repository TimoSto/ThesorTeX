package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/TimoSto/ThesorTeX/pkg/backend/handler_chain"
	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"github.com/TimoSto/ThesorTeX/pkg/backend/pathbuilder"
	"github.com/TimoSto/ThesorTeX/services/app/internal/config"
	"github.com/TimoSto/ThesorTeX/services/app/internal/domain/projects"
	"github.com/TimoSto/ThesorTeX/services/app/internal/filesystem/fake"
	"github.com/TimoSto/ThesorTeX/services/app/internal/handlers"
)

func main() {
	log.Info("Starting the e2e app...")
	log.Info("Version: %s", config.Version)

	//pathbuilder.Init()

	mux := http.NewServeMux()

	chain := handler_chain.CreateHandlerChain()

	fs := fake.FileSystem{}

	config.Cfg = config.Config{
		ProjectsDir: "projects",
	}

	log.Info("Creating example project...")
	_, err := projects.CreateProject("example", &fs, config.Cfg)
	if err != nil {
		log.Error("unexpected error creating the example project: %v", err)
		os.Exit(1)
	}
	log.Info("Created example project under %s", pathbuilder.GetProjectPath(config.Cfg.ProjectsDir, "example"))

	handlers.RegisterAppHandlers(mux, &fs)

	err = http.ListenAndServe(fmt.Sprintf("localhost:%s", "8440"), chain.Then(mux))
	if err != nil {
		log.Error("unexpected error starting server: %v", err)
		os.Exit(1)
	}
}
