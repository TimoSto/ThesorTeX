package caching

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

var (
	CacheTimeHTML   = 1 * time.Minute
	CacheTimeImages = 5 * time.Minute
	CacheTimeJSON   = 5 * time.Minute
	CacheTimeAssets = 24 * time.Hour
)

func ApplyCacheTimes(r *http.Request, w *http.ResponseWriter) {

	ct := r.Header.Get("Content-Type")
	pos := strings.LastIndex(r.URL.Path, ".")

	if pos >= 0 {
		ct = r.URL.Path[pos+1 : len(r.URL.Path)]
	}

	var cache time.Duration

	fmt.Println(r.URL.Path)

	if ct == "html" || r.URL.Path == "/" {
		cache = CacheTimeHTML
	} else if ct == "css" || ct == "js" {
		cache = CacheTimeAssets
	} else if ct == "png" {
		cache = CacheTimeImages
	} else if ct == "json" {
		cache = CacheTimeJSON
	}

	rw := *w

	rw.Header().Set("Cache-Control", fmt.Sprintf("max-age=%v", int(cache.Seconds())))
}
