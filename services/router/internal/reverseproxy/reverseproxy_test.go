package reverseproxy_test

import (
	"github.com/TimoSto/ThesorTeX/services/router/internal/reverseproxy"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegisterProxyHandler(t *testing.T) {
	// TODO: test this
	for _, tc := range []struct {
		name     string
		callsTo1 int
		callsTo2 int
	}{
		{
			name:     "call server 1 once",
			callsTo1: 1,
			callsTo2: 0,
		},
		{
			name:     "call server 2 once",
			callsTo1: 0,
			callsTo2: 1,
		},
		{
			name:     "call both once",
			callsTo1: 1,
			callsTo2: 1,
		},
		{
			name:     "call both twice",
			callsTo1: 2,
			callsTo2: 2,
		},
		{
			name:     "call server 1 once and server 2 twice",
			callsTo1: 1,
			callsTo2: 2,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			m := http.NewServeMux()

			called1 := 0
			h1 := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				called1++
			})
			s1 := httptest.NewServer(h1)
			defer s1.Close()

			err := reverseproxy.RegisterProxyHandler(m, "/", s1.URL)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			called2 := 0
			h2 := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				called2++
			})
			s2 := httptest.NewServer(h2)
			defer s2.Close()

			err = reverseproxy.RegisterProxyHandler(m, "/s2/", s2.URL)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			for i := 0; i < tc.callsTo1; i++ {
				r := httptest.NewRequest("GET", "/", nil)
				rr := httptest.NewRecorder()
				m.ServeHTTP(rr, r)
			}

			for i := 0; i < tc.callsTo2; i++ {
				r := httptest.NewRequest("GET", "/s2/", nil)
				rr := httptest.NewRecorder()
				m.ServeHTTP(rr, r)
			}

			if called1 != tc.callsTo1 {
				t.Errorf("expected server 1 to have been called %v times, but got %v", tc.callsTo1, called1)
			}

			if called2 != tc.callsTo2 {
				t.Errorf("expected server 2 to have been called %v times, but got %v", tc.callsTo2, called2)
			}
		})
	}
}
