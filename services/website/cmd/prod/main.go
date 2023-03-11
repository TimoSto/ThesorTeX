package main

import (
	logD "log"
	"net/http"

	"github.com/TimoSto/ThesorTeX/pkg/backend/handler_chain"
	"github.com/TimoSto/ThesorTeX/pkg/backend/lambda"
	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"github.com/TimoSto/ThesorTeX/pkg/backend/s3"
	"github.com/TimoSto/ThesorTeX/services/website/internal/buckethandler"
	"github.com/TimoSto/ThesorTeX/services/website/internal/handlers"
)

func main() {
	log.Info("Start lambda")

	mux := http.NewServeMux()

	s3Client, err := s3.CreateS3Client()
	if err != nil {
		logD.Fatal(err)
	}

	s3Handler := buckethandler.New(s3Client)

	handlers.RegisterWebsiteHandlers(mux, &s3Handler)

	chain := handler_chain.CreateHandlerChain()

	lambda.Start(chain.Then(mux))
}
