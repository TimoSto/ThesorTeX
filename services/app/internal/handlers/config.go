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
				log.Error(fmt.Printf("Mershaling config data: %v", err))
				return
			}

			_, err = w.Write(data)
			if err != nil {
				log.Error(fmt.Printf("Sending config data to client: %v", err))
				return
			}

		case http.MethodPost:
			var data conf.Config
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&data)
			if err != nil {
				log.Error(fmt.Sprintf("Saving config: %v", err))
				return
			}
			err = conf.WriteConfig(data)
			if err != nil {
				log.Error(fmt.Printf("Saving config: %v", err))
				return
			}

		default:
			w.WriteHeader(http.StatusBadRequest)
		}
	}

	return http.HandlerFunc(fn)
}
