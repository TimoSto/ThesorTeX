package handlers

import (
	"net/http"

	"github.com/TimoSto/ThesorTeX/services/website/internal/buckethandler"
	"github.com/TimoSto/ThesorTeX/services/website/internal/handlers/assets"
	"github.com/TimoSto/ThesorTeX/services/website/internal/handlers/documentations"
	"github.com/TimoSto/ThesorTeX/services/website/internal/handlers/examples"
	"github.com/TimoSto/ThesorTeX/services/website/internal/handlers/templates"
	"github.com/TimoSto/ThesorTeX/services/website/internal/handlers/versions"
)

func RegisterWebsiteHandlers(mux *http.ServeMux, dev bool, bucket buckethandler.BucketHandler) {
	mux.HandleFunc("/", assets.HandleAssets())

	mux.HandleFunc("/templates/thesis", templates.HandleThesisTemplate())

	mux.HandleFunc("/templates/cv", templates.HandleCVTemplate())

	mux.HandleFunc("/versions", versions.HandleVersions(dev, &bucket))

	mux.HandleFunc("/documentation", documentations.HandleDocumentations())

	mux.HandleFunc("/example", examples.HandleExamples())
}
