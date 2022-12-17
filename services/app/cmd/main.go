package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/TimoSto/ThesorTeX/pkg/handler_chain"
	"github.com/TimoSto/ThesorTeX/pkg/log"
	"github.com/TimoSto/ThesorTeX/services/app/internal/conf"
	"github.com/TimoSto/ThesorTeX/services/app/internal/handlers"
)

func main() {
	log.Info("Starting local app...")

	conf.ReadConfig()

	mux := http.NewServeMux()

	chain := handler_chain.CreateHandlerChain()

	//assets-handler bei beiden gleich, nur andere ignore-parameter
	//assetConf := assets.AssetConf{
	//	Ignore: []string{
	//		"/auth.html",
	//	},
	//	MapRootTo: "/app.html",
	//}
	//assets.Register(mux, assetConf)

	handlers.Register(mux, conf.GetConfig())

	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%s", conf.GetConfig().Port), chain.Then(mux))
		if err != nil {
			log.Error("unexpected error starting server: %v", err)
			os.Exit(1)
		}
	}()

	log.Info("App running under http://localhost:8448")

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	log.Info("waiting for exit...")

	sig := <-sigs

	log.Info("received %v", sig)
}
