package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/TimoSto/ThesorTeX/pkg/handler_chain"
	"github.com/TimoSto/ThesorTeX/services/project_management/conf"
	"github.com/TimoSto/ThesorTeX/webclient/assets"
)

func main() {
	fmt.Println("Starting local app...")

	configObj := conf.GetConfig()

	mux := http.NewServeMux()

	chain := handler_chain.CreateHandlerChain()

	//assets-handler bei beiden gleich, nur andere ignore-parameter
	assets.Register(mux)

	go http.ListenAndServe(fmt.Sprintf(":%s", configObj.Port), chain.Then(mux))

	fmt.Println("App running under http://localhost:8448")

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("waiting for exit...")

	sig := <-sigs

	fmt.Println("received", sig)
}
