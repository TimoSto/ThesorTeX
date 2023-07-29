package main

import (
	"fmt"
	"github.com/TimoSto/ThesorTeX/pkg/backend/http/server"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"syscall"

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
	log.Setup(false, "INFO")

	log.Info("Version: %s", config.Version)

	log.Info("Starting the local app...")

	pathbuilder.Init()

	mux := http.NewServeMux()

	chain := handler_chain.CreateHandlerChain()

	fs := local.FileSystem{}

	_, err := config.ReadConfig()
	if err != nil {
		log.Error("unexpected error reading the config: %v", err)
		os.Exit(1)
	}
	log.Setup(false, config.Cfg.LogLevel)

	updatechecker.CheckUpdateAvailable()

	exists, err := fs.CheckDirectoryExists(config.Cfg.ProjectsDir)
	if err != nil {
		log.Error("unexpected error reading the projects dir: %v", err)
		os.Exit(1)
	}

	if !exists {
		log.Info("Creating example project...")
		_, err = projects.CreateProject("example", &fs, config.Cfg)
		if err != nil {
			log.Error("unexpected error creating the example project: %v", err)
			os.Exit(1)
		}
		log.Info("Created example project under %s", pathbuilder.GetProjectPath(config.Cfg.ProjectsDir, "example"))
	}

	handlers.RegisterAppHandlers(mux, &fs)

	finished := make(chan bool, 1)

	p, err := server.StartServer(config.Cfg.Port, chain.Then(mux), finished)
	if err != nil {
		log.Fatal("could not start server: %v", err)
	}

	addr := fmt.Sprintf("http://localhost:%s", p)
	log.Info("local server is running under %s", addr)

	if config.Cfg.OpenBrowser {
		openBrowser(addr)
	}

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	log.Info("waiting for exit...")

	sig := <-sigs

	log.Info("received %v", sig)

	finished <- true
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
