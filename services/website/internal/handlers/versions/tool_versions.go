package versions

import (
	"encoding/json"
	"net/http"

	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"github.com/TimoSto/ThesorTeX/services/website/internal/versions"
)

func HandleToolVersions(dev bool) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		toolVersions := versions.GetToolVersions(dev)

		data, err := json.Marshal(toolVersions)

		if err != nil {
			log.Error("could not serialize versions: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(data)
	}

	return fn
}
