package backend

import (
	"github.com/TimoSto/ThesorTeX/pkg/backend/handler_chain"
	"github.com/TimoSto/ThesorTeX/services/router/internal/config"
	"github.com/TimoSto/ThesorTeX/services/router/internal/reverseproxy"
	"net/http"
)

func StartBackend() error {
	ms, err := config.GetMappings()
	if err != nil {
		return err
	}

	mux := http.NewServeMux()

	for _, m := range ms {
		err := reverseproxy.RegisterProxyHandler(mux, m.Route, m.TargetURL)
		if err != nil {
			return err
		}
	}

	chain := handler_chain.CreateHandlerChain()

	err = http.ListenAndServe(":8446", chain.Then(mux))
	if err != nil {
		return err
	}

	return nil
}
