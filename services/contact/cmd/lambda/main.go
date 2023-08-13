package main

import (
	"fmt"
	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"github.com/TimoSto/ThesorTeX/services/contact/internal/backend"
)

func main() {
	log.Setup(true, "DEBUG")

	fmt.Println("hallo")

	log.Info("Starting the aws contact application...")

	cfg := backend.Config{
		IsProd: true,
	}

	err := backend.StartApp(cfg)
	if err != nil {
		log.Fatal("could not start backend: %v", err)
	}
}
