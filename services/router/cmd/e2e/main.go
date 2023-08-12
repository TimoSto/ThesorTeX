package main

import (
	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"github.com/TimoSto/ThesorTeX/services/router/internal/backend"
	"os"
)

func main() {

	log.Setup(false, "DEBUG")

	sigs := make(chan os.Signal, 1)

	cfg := backend.Config{
		IsProd: false,
	}

	err := backend.StartBackend(cfg)
	if err != nil {
		log.Fatal("could not start backend: %v", err)
	}

	<-sigs
}
