package server_test

import (
	"github.com/TimoSto/ThesorTeX/pkg/backend/http/server"
	"net/http"
	"testing"
)

func TestNewServer_fixedPort(t *testing.T) {
	h := func(w http.ResponseWriter, r *http.Request) {

	}

	p, err := server.StartServer("8081", http.HandlerFunc(h))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if want := "8081"; want != p {
		t.Errorf("expected port %s but got %s", want, p)
	}
}

func TestNewServer_zeroPort(t *testing.T) {
	h := func(w http.ResponseWriter, r *http.Request) {

	}

	p, err := server.StartServer("0", http.HandlerFunc(h))
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if "0" == p {
		t.Error("expected port different than 0")
	}
}
