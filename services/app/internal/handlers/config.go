package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TimoSto/ThesorTeX/pkg/log"
	"github.com/TimoSto/ThesorTeX/services/app/conf"
)

func HandleConfig() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {

		case http.MethodGet:
			data, err := json.Marshal(conf.GetConfig())
			if err != nil {
				log.Error(fmt.Printf("MArshaling config data: %v", err))
				return
			}

			w.Write(data)

		case http.MethodPost:
			var data conf.Config
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&data)
			if err != nil {
				log.Error(fmt.Sprintf("Saving config: %v", err))
				return
			}
			conf.WriteConfig(data)
		}
	}

	return http.HandlerFunc(fn)
}
