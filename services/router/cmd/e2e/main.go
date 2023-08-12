package main

import (
	"github.com/TimoSto/ThesorTeX/services/router/internal/backend"
	"log"
	"os"
)

func main() {

	sigs := make(chan os.Signal, 1)

	err := backend.StartBackend()
	if err != nil {
		log.Fatalf("could not start backend: %v", err)
	}

	<-sigs
}
