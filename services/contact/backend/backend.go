package backend

import (
	"fmt"
	"github.com/TimoSto/ThesorTeX/pkg/backend/aws/apigateway"
	"github.com/TimoSto/ThesorTeX/pkg/backend/handler_chain"
	"net/http"
)

type Config struct {
	IsProd bool
}

func StartApp(cfg Config) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("got req")
	})

	chain := handler_chain.CreateHandlerChain()

	if cfg.IsProd {
		apigateway.StartLambda(chain.Then(mux))
	} else {
		err := http.ListenAndServe(":8447", chain.Then(mux))
		if err != nil {
			return err
		}
	}

	return nil
}
