package main

import (
	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"github.com/TimoSto/ThesorTeX/services/app/internal/config"
)

func main() {
	log.Info("Starting the local app...")
	log.Info("Version: %s", config.Version)
}
