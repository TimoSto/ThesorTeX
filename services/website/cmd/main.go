package main

import (
	"fmt"
	"net/http"

	"github.com/TimoSto/ThesorTeX/pkg/backend/handler_chain"
	"github.com/TimoSto/ThesorTeX/pkg/backend/lambda"
	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
)

// Handler is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func Handler(w http.ResponseWriter, r *http.Request) {
	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Info("got Request %s", r.URL.Path)

	w.Write([]byte(fmt.Sprintf("hallo %s", r.URL.Path)))
}

func main() {
	log.Info("Start lambda")

	mux := http.NewServeMux()

	mux.HandleFunc("/", Handler)

	chain := handler_chain.CreateHandlerChain()

	lambda.Start(chain.Then(mux))
}
