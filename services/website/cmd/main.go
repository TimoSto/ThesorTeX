package main

import (
	"fmt"
	logD "log"
	"net/http"
	"os"

	"github.com/TimoSto/ThesorTeX/pkg/backend/handler_chain"
	"github.com/TimoSto/ThesorTeX/pkg/backend/lambda"
	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"github.com/TimoSto/ThesorTeX/pkg/backend/s3"
	"github.com/TimoSto/ThesorTeX/services/website/internal/handlers"
)

func main() {
	log.Info("Start lambda")

	mux := http.NewServeMux()

	dev := os.Getenv("DEV")

	s3Client, err := s3.CreateS3Client()
	if err != nil {
		logD.Fatal(err)
	}

	handlers.RegisterWebsiteHandlers(mux, dev == "true", s3Client)

	chain := handler_chain.CreateHandlerChain()

	if dev == "true" {
		fmt.Println("run local")
		err := http.ListenAndServe("localhost:8449", chain.Then(mux))
		fmt.Println(err)
	} else {
		lambda.Start(chain.Then(mux))
	}
}
