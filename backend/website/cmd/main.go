package main

import (
	"net/http"
	"os"

	"github.com/TimoSto/ThesorTeX/backend/pkg/handler_chain"
	"github.com/TimoSto/ThesorTeX/backend/pkg/log"
	"github.com/TimoSto/ThesorTeX/backend/website/internal/static_content"
)

func main() {

	mux := http.NewServeMux()

	chain := handler_chain.CreateHandlerChain()

	static_content.Register(mux)

	if os.Getenv("RUN_LOCAL") == "true" {
		log.Info("Running locally on http://localhost:8558 ")
		err := http.ListenAndServe(":8558", chain.Then(mux))
		if err != nil {
			log.Error("unexpected error starting the server: %v", err)
		}
	}

}
