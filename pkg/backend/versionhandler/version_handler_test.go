package versionhandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetRootHandler(t *testing.T) {
	for _, tc := range []struct {
		method        string
		version       string
		expectedCode  int
		expectedBody  versionRes
		expectedCache string
	}{
		{
			method:        "POST",
			version:       "1.2.3",
			expectedCode:  http.StatusMethodNotAllowed,
			expectedCache: fmt.Sprintf("max-age=86400"),
		},
		{
			method:        "PUT",
			version:       "1.2.3",
			expectedCode:  http.StatusMethodNotAllowed,
			expectedCache: fmt.Sprintf("max-age=86400"),
		},
		{
			method:        "DELETE",
			version:       "1.2.3",
			expectedCode:  http.StatusMethodNotAllowed,
			expectedCache: fmt.Sprintf("max-age=86400"),
		},
		{
			method:        "GET",
			version:       "1.2.3",
			expectedCode:  http.StatusOK,
			expectedBody:  versionRes{Version: "1.2.3"},
			expectedCache: fmt.Sprintf("max-age=86400"),
		},
	} {
		t.Run(tc.method, func(t *testing.T) {
			req, err := http.NewRequest(tc.method, "/favicon", nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(GetRootHandler(tc.version))

			handler.ServeHTTP(rr, req)

			if rr.Code != tc.expectedCode {
				t.Errorf("expected status code %v, but got %v", tc.expectedCode, rr.Code)
			}

			var gotVersion versionRes
			json.Unmarshal(rr.Body.Bytes(), &gotVersion)
			if diff := cmp.Diff(gotVersion, tc.expectedBody); diff != "" {
				t.Errorf("unexpected diff in version: %s", diff)
			}

			if got := rr.Header().Get("Cache-Control"); got != tc.expectedCache {
				t.Errorf("expected cache %s, but got %s", tc.expectedCache, got)
			}
		})
	}
}
