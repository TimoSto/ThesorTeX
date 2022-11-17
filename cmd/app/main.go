package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/TimoSto/ThesorTeX/pkg/handler_chain"
	"github.com/TimoSto/ThesorTeX/services/project_management/register"
)

func main() {
	fmt.Println("Starting local app...")

	mux := http.NewServeMux()

	chain := handler_chain.CreateHandlerChain()

	register.Register(mux)

	http.ListenAndServe(":8448", chain.Then(mux))

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("waiting for exit...")

	sig := <-sigs

	fmt.Println("received", sig)
}
