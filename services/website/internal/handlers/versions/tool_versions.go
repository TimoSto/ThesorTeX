package versions

import (
	"encoding/json"
	"net/http"

	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"github.com/TimoSto/ThesorTeX/services/website/internal/buckethandler"
	"github.com/TimoSto/ThesorTeX/services/website/internal/versions"
)

func HandleVersions(dev bool, bucket *buckethandler.BucketHandler) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		allVersions, err := versions.GetVersions(dev, bucket)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data, err := json.Marshal(allVersions)

		if err != nil {
			log.Error("could not serialize versions: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(data)
	}

	return fn
}
