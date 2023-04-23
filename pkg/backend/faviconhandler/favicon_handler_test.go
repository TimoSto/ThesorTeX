package faviconhandler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetFaviconHandler(t *testing.T) {

	for _, tc := range []struct {
		method       string
		expectedCode int
		expectedBody []byte
	}{
		{
			method:       "POST",
			expectedCode: http.StatusMethodNotAllowed,
		},
		{
			method:       "PUT",
			expectedCode: http.StatusMethodNotAllowed,
		},
		{
			method:       "DELETE",
			expectedCode: http.StatusMethodNotAllowed,
		},
		{
			method:       "GET",
			expectedCode: http.StatusOK,
			expectedBody: favicon,
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
		})
	}

}
