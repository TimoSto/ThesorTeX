package backend

import (
	"github.com/TimoSto/ThesorTeX/pkg/backend/aws/apigateway"
	"github.com/TimoSto/ThesorTeX/pkg/backend/handler_chain"
	"github.com/TimoSto/ThesorTeX/services/contact/internal/handlers"
	"net/http"
)

type Config struct {
	IsProd bool
	Mux    *http.ServeMux
}

func StartApp(cfg Config) error {
	if cfg.Mux == nil {
		cfg.Mux = http.NewServeMux()
	}

	handlers.RegisterHandlers(cfg.Mux)

	chain := handler_chain.CreateHandlerChain()

	if cfg.IsProd {
		apigateway.StartLambda(chain.Then(cfg.Mux))
	} else {
		err := http.ListenAndServe(":8447", chain.Then(cfg.Mux))
		if err != nil {
			return err
		}
	}

	return nil
}
