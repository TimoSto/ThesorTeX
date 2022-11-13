package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/TimoSto/ThesorTeX/internal/backend"
)

func main() {
	fmt.Println("Start")

	backend.Start()

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("waiting for exit...")

	sig := <-sigs

	fmt.Println("received", sig)
}
