package main

import (
	"net/http"
	"os"

	"github.com/TimoSto/ThesorTeX/pkg/handler_chain"
	"github.com/TimoSto/ThesorTeX/pkg/log"
)

func main() {

	mux := http.NewServeMux()

	chain := handler_chain.CreateHandlerChain()

	if os.Getenv("RUN_LOCAL") == "true" {
		log.Info("Running locally on http://localhost:8558 ")
		log.Error(http.ListenAndServe(":8558", chain.Then(mux)))
	}

}
