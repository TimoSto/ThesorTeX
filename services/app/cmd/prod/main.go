package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/TimoSto/ThesorTeX/pkg/backend/handler_chain"
	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"github.com/TimoSto/ThesorTeX/pkg/backend/pathbuilder"
	"github.com/TimoSto/ThesorTeX/services/app/internal/config"
	"github.com/TimoSto/ThesorTeX/services/app/internal/domain/projects"
	"github.com/TimoSto/ThesorTeX/services/app/internal/filesystem/local"
	"github.com/TimoSto/ThesorTeX/services/app/internal/handlers"
)

func main() {
	log.Info("Starting the local app...")
	log.Info("Version: %s", config.Version)

	pathbuilder.Init()

	mux := http.NewServeMux()

	chain := handler_chain.CreateHandlerChain()

	fs := local.FileSystem{}

	cfg, err := config.ReadConfig()
	if err != nil {
		log.Error("unexpected error reading the config: %v", err)
		os.Exit(1)
	}

	exists, err := fs.CheckDirectoryExists(cfg.ProjectsDir)
	if err != nil {
		log.Error("unexpected error reading the projects dir: %v", err)
		os.Exit(1)
	}

	if !exists {
		log.Info("Creating example project...")
		_, err = projects.CreateProject("example", &fs, cfg)
		if err != nil {
			log.Error("unexpected error creating the example project: %v", err)
			os.Exit(1)
		}
		log.Info("Created example project under %s", pathbuilder.GetProjectPath(cfg.ProjectsDir, "example"))
	}

	handlers.RegisterAppHandlers(mux, &fs, cfg)

	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), chain.Then(mux))
		if err != nil {
			log.Error("unexpected error starting server: %v", err)
			os.Exit(1)
		}
	}()

	//TODO: get actual url icase of port 0

	log.Info("App running under http://localhost:%s", cfg.Port)

	if cfg.OpenBrowser {
		openBrowser(fmt.Sprintf("http://localhost:%s", cfg.Port))
	}

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	log.Info("waiting for exit...")

	sig := <-sigs

	log.Info("received %v", sig)
}

func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Error("Could not open browser: %v", err)
	}
}
