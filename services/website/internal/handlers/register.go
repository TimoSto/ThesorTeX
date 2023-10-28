package handlers

import (
	"github.com/TimoSto/ThesorTeX/services/website/internal/dsgvo/dataProtectionDeclaration"
	"net/http"

	"github.com/TimoSto/ThesorTeX/pkg/backend/faviconhandler"
	"github.com/TimoSto/ThesorTeX/pkg/backend/versionhandler"
	"github.com/TimoSto/ThesorTeX/services/website/internal/buckethandler"
	"github.com/TimoSto/ThesorTeX/services/website/internal/config"
	"github.com/TimoSto/ThesorTeX/services/website/internal/handlers/assets"
	"github.com/TimoSto/ThesorTeX/services/website/internal/handlers/documentations"
	"github.com/TimoSto/ThesorTeX/services/website/internal/handlers/templates"
	"github.com/TimoSto/ThesorTeX/services/website/internal/handlers/versions"
)

func RegisterWebsiteHandlers(mux *http.ServeMux, bucket buckethandler.BucketHandler) {
	mux.HandleFunc("/", assets.HandleAssets())

	mux.HandleFunc("/version", versionhandler.GetRootHandler(config.Version))

	mux.HandleFunc("/favicon", faviconhandler.GetFaviconHandler())

	mux.HandleFunc("/templates/thesis", templates.HandleThesisTemplate())

	mux.HandleFunc("/templates/cv", templates.HandleCVTemplate())

	mux.HandleFunc("/versions", versions.HandleVersions(&bucket))

	mux.HandleFunc("/documentation", documentations.HandleDocumentations())

	mux.HandleFunc("/documentation/images/", documentations.HandleDocumentationImages())

	mux.HandleFunc("/dsgvo", dataProtectionDeclaration.HandleDSGVO())
}
