package backend

import (
	"github.com/TimoSto/ThesorTeX/pkg/backend/aws/apigateway"
	"github.com/TimoSto/ThesorTeX/pkg/backend/aws/dynamodb"
	"github.com/TimoSto/ThesorTeX/pkg/backend/handler_chain"
	"github.com/TimoSto/ThesorTeX/services/contact/internal/feedback"
	"github.com/TimoSto/ThesorTeX/services/contact/internal/handlers"
	"github.com/aws/aws-sdk-go-v2/config"
	"net/http"
)

type Config struct {
	IsProd     bool
	Mux        *http.ServeMux
	DynamoOpts config.LoadOptionsFunc
}

func StartApp(cfg Config) error {
	if cfg.Mux == nil {
		cfg.Mux = http.NewServeMux()
	}

	client, err := dynamodb.GetDynamoClient(cfg.DynamoOpts)
	if err != nil {
		return err
	}

	store := feedback.New(client)

	handlers.RegisterHandlers(cfg.Mux, store)

	chain := handler_chain.CreateHandlerChain()

	if cfg.IsProd {
		apigateway.StartLambda(chain.Then(chain.Then(cfg.Mux)))
	} else {
		err := http.ListenAndServe(":8447", chain.Then(cfg.Mux))
		if err != nil {
			return err
		}
	}

	return nil
}
