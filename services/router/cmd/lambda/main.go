package main

import (
	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"github.com/TimoSto/ThesorTeX/services/router/internal/backend"
)

func main() {
	log.Setup(true, "DEBUG")
	cfg := backend.Config{
		IsProd: true,
	}

	err := backend.StartBackend(cfg)
	if err != nil {
		log.Fatal("could not start backend: %v", err)
	}
}
