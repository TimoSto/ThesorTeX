package faviconhandler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetFaviconHandler(t *testing.T) {

	for _, tc := range []struct {
		method        string
		expectedCode  int
		expectedBody  []byte
		expectedCache string
	}{
		{
			method:        "POST",
			expectedCode:  http.StatusMethodNotAllowed,
			expectedCache: fmt.Sprintf("max-age=86400"),
		},
		{
			method:        "PUT",
			expectedCode:  http.StatusMethodNotAllowed,
			expectedCache: fmt.Sprintf("max-age=86400"),
		},
		{
			method:        "DELETE",
			expectedCode:  http.StatusMethodNotAllowed,
			expectedCache: fmt.Sprintf("max-age=86400"),
		},
		{
			method:        "GET",
			expectedCode:  http.StatusOK,
			expectedBody:  favicon,
			expectedCache: fmt.Sprintf("max-age=86400"),
		},
	} {
		t.Run(tc.method, func(t *testing.T) {
			req, err := http.NewRequest(tc.method, "/favicon", nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(GetFaviconHandler())

			handler.ServeHTTP(rr, req)

			if rr.Code != tc.expectedCode {
				t.Errorf("expected status code %v, but got %v", tc.expectedCode, rr.Code)
			}

			if rr.Body.String() != string(tc.expectedBody) {
				t.Errorf("expected body %v, but got %v", string(tc.expectedBody), rr.Body.String())
			}

			if got := rr.Header().Get("Cache-Control"); got != tc.expectedCache {
				t.Errorf("expected cache %s, but got %s", tc.expectedCache, got)
			}
		})
	}

}
