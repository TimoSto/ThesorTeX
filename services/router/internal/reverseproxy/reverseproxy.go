package reverseproxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func RegisterProxyHandler(mux *http.ServeMux, path, serverURL string) error {
	origin, err := url.Parse(serverURL)
	if err != nil {
		return err
	}
	proxy := httputil.NewSingleHostReverseProxy(origin)

	mux.Handle(path, proxy)

	return nil
}
