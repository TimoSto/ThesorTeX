package handler_chain

import (
	"net/http"

	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"github.com/justinas/alice"
)

func CreateHandlerChain() alice.Chain {
	return alice.New(
		LogRequest(),
		//TODO: add headers
		//TODO: unit test
	)
}

func LogRequest() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			log.Info("Handling request: %s", req.URL.Path)

			next.ServeHTTP(rw, req)
		})
	}
}
