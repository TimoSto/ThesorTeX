package versions

import (
	"encoding/json"
	"net/http"

	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"github.com/TimoSto/ThesorTeX/services/website/internal/versions"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func HandleToolVersions(dev bool, s3Client *s3.Client) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		toolVersions, err := versions.GetToolVersions(dev, s3Client)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

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
