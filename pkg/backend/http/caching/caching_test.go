package caching

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func MockHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hallo")
	ApplyCacheTimes(r, &w)

	fmt.Println(w.Header().Get("Cache-Control"))

	w.WriteHeader(200)
}

func TestApplyCaching(t *testing.T) {
	for _, tc := range []struct {
		file          string
		expectedCache string
	}{
		{
			file:          "text.html",
			expectedCache: "max-age=60",
		},
		{
			file:          "text.css",
			expectedCache: "max-age=86400",
		},
		{
			file:          "text.js",
			expectedCache: "max-age=86400",
		},
		{
			file:          "image.png",
			expectedCache: "max-age=300",
		},
		{
			file:          "application.json",
			expectedCache: "max-age=300",
		},
	} {
		t.Run(tc.file, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", fmt.Sprintf("https://test.com/%s", tc.file), nil)

			handler := http.HandlerFunc(MockHandler)

			handler.ServeHTTP(recorder, req)

			got := recorder.Result().Header.Get("Cache-Control")

			if got != tc.expectedCache {
				t.Errorf("got %s, epected %s", got, tc.expectedCache)
			}
		})
	}

}
