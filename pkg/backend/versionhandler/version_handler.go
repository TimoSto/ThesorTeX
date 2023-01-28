package versionhandler

import (
	"encoding/json"
	"net/http"

	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
)

type versionRes struct {
	Version string
}

func GetRootHandler(version string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		data, err := json.Marshal(versionRes{Version: version})
		if err != nil {
			log.Error("Marshaling version data: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(data)
	}

	return fn
}
