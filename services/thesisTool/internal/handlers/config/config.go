package config

import (
	"encoding/json"
	"net/http"

	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"github.com/TimoSto/ThesorTeX/services/thesisTool/internal/config"
)

func HandleConfigGet() func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		data, err := json.Marshal(config.Cfg)
		if err != nil {
			log.Error("could not serialize config data: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(data)
	}

	return fn
}

//IDEA: cfg not as variable but as property in config package

func HandleConfigSave() func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var gotConfig config.Config

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&gotConfig)
		if err != nil {
			log.Error("cloud not deserialize config data to save: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = config.SaveConfig(gotConfig)
		if err != nil {
			log.Error("could not save config: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

	return fn
}
