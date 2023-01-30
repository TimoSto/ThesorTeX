package config

import (
	"encoding/json"
	"net/http"

	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"github.com/TimoSto/ThesorTeX/services/app/internal/config"
)

func HandleConfigGet(cfg config.Config) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		data, err := json.Marshal(cfg)
		if err != nil {
			log.Error("could not serialize config data: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(data)
	}

	return fn
}
