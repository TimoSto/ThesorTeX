package main

import (
	"net/http"

	"github.com/TimoSto/ThesorTeX/pkg/backend/aws/apigateway"
	"github.com/TimoSto/ThesorTeX/pkg/backend/aws/s3"
	"github.com/TimoSto/ThesorTeX/pkg/backend/handler_chain"
	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"github.com/TimoSto/ThesorTeX/services/website/internal/buckethandler"
	"github.com/TimoSto/ThesorTeX/services/website/internal/handlers"
)

func main() {
	log.Setup(true)

	log.Info("Start lambda")

	mux := http.NewServeMux()

	s3Client, err := s3.CreateS3Client()
	if err != nil {
		log.Fatal("could not create s3 client: %v", err)
	}

	s3Handler := buckethandler.New(s3Client)

	handlers.RegisterWebsiteHandlers(mux, &s3Handler)

	chain := handler_chain.CreateHandlerChain()

	apigateway.StartLambda(chain.Then(mux))
}
