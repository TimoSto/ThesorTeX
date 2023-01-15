package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/TimoSto/ThesorTeX/pkg/backend/handler_chain"
	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"github.com/TimoSto/ThesorTeX/services/app/internal/config"
	"github.com/TimoSto/ThesorTeX/services/app/internal/handlers"
)

func main() {
	log.Info("Starting the local app...")
	log.Info("Version: %s", config.Version)

	mux := http.NewServeMux()

	chain := handler_chain.CreateHandlerChain()

	handlers.RegisterAppHandlers(mux)

	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%s", "8448"), chain.Then(mux))
		if err != nil {
			log.Error("unexpected error starting server: %v", err)
			os.Exit(1)
		}
	}()

	log.Info("App running under http://localhost:8448")

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	log.Info("waiting for exit...")

	sig := <-sigs

	log.Info("received %v", sig)
}
