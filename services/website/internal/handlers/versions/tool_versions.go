package versions

import (
	"encoding/json"
	"net/http"

	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
)

func HandleToolVersions(dev bool) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		var versions []string

		if dev {
			versions = []string{
				"0.0.1",
				"0.0.2",
				"1.0.0",
			}
		}

		data, err := json.Marshal(versions)

		if err != nil {
			log.Error("could not serialize versions: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(data)
	}

	return fn
}
