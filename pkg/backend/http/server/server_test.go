package server_test

import (
	"github.com/TimoSto/ThesorTeX/pkg/backend/http/server"
	"net/http"
	"testing"
)

func TestNewServer_fixedPort(t *testing.T) {
	h := func(w http.ResponseWriter, r *http.Request) {

	}

	s := server.NewServer("8081", http.HandlerFunc(h), make(chan bool, 1))

	if want := "8081"; want != s.Port() {
		t.Errorf("expected port %s but got %s", want, s.Port())
	}
}

func TestNewServer_zeroPort(t *testing.T) {
	h := func(w http.ResponseWriter, r *http.Request) {

	}

	s := server.NewServer("0", http.HandlerFunc(h), make(chan bool, 1))

	if "0" == s.Port() {
		t.Error("expected port different than 0")
	}
}
