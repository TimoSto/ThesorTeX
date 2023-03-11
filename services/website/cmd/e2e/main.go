package main

import (
	"fmt"
	"net/http"

	"github.com/TimoSto/ThesorTeX/pkg/backend/handler_chain"
	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"github.com/TimoSto/ThesorTeX/services/website/internal/buckethandler"
	"github.com/TimoSto/ThesorTeX/services/website/internal/handlers"
)

func main() {
	log.Info("Start fake website")

	bucket := buckethandler.FakeBucketHandler{}

	bucket.SetItem("thesisTool/v1.1.0", "test-2")
	bucket.SetItem("thesisTool/v1.0.0", "test-1")
	bucket.SetItem("thesisTemplate/v1.0.0", "test-1")
	bucket.SetItem("cvTemplate/v1.0.0", "test-1")

	mux := http.NewServeMux()

	handlers.RegisterWebsiteHandlers(mux, &bucket)

	chain := handler_chain.CreateHandlerChain()

	err := http.ListenAndServe("localhost:8449", chain.Then(mux))
	fmt.Println(err)

}
