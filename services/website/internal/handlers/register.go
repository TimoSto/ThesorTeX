package handlers

import (
	"net/http"

	"github.com/TimoSto/ThesorTeX/services/website/internal/handlers/assets"
	"github.com/TimoSto/ThesorTeX/services/website/internal/handlers/documentations"
	"github.com/TimoSto/ThesorTeX/services/website/internal/handlers/examples"
	"github.com/TimoSto/ThesorTeX/services/website/internal/handlers/templates"
	"github.com/TimoSto/ThesorTeX/services/website/internal/handlers/versions"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func RegisterWebsiteHandlers(mux *http.ServeMux, dev bool, s3Client *s3.Client) {
	mux.HandleFunc("/", assets.HandleAssets())

	mux.HandleFunc("/templates/thesis", templates.HandleThesisTemplate())

	mux.HandleFunc("/templates/cv", templates.HandleCVTemplate())

	mux.HandleFunc("/versions/tool", versions.HandleToolVersions(dev, s3Client))

	mux.HandleFunc("/documentation", documentations.HandleDocumentations())

	mux.HandleFunc("/example", examples.HandleExamples())
}
