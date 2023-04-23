package versionhandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
)

type versionRes struct {
	Version string
}

func GetRootHandler(version string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", fmt.Sprintf("max-age=%v", int(24*time.Hour/time.Second)))

		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
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
