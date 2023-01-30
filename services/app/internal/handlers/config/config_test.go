package config

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/TimoSto/ThesorTeX/services/app/internal/config"
)

func TestHandleConfigGet_POST(t *testing.T) {
	req, err := http.NewRequest("POST", "/getAllProjects", nil)
	if err != nil {
		t.Fatal(err)
	}

	cfg := config.Config{
		ProjectsDir: "projects",
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleConfigGet(cfg))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("expected status code %v, but got %v", http.StatusMethodNotAllowed, status)
	}
}

func TestHandleConfigGet_GET(t *testing.T) {
	req, err := http.NewRequest("GET", "/getAllProjects", nil)
	if err != nil {
		t.Fatal(err)
	}

	cfg := config.Config{
		ProjectsDir: "projects",
		Port:        "5678",
		OpenBrowser: true,
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleConfigGet(cfg))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status code %v, but got %v", http.StatusOK, status)
	}

	var gotConfig config.Config

	decoder := json.NewDecoder(rr.Body)
	err = decoder.Decode(&gotConfig)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(cfg, gotConfig) {
		t.Errorf("expected %v, got %v", cfg, gotConfig)
	}
}

func TestHandleConfigPost_GET(t *testing.T) {

	req, err := http.NewRequest("GET", "/getAllProjects", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleConfigSave())

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("expected status code %v, but got %v", http.StatusMethodNotAllowed, status)
	}
}
func TestHandleConfigPost_POST(t *testing.T) {
	cfg := config.Config{
		ProjectsDir: "projects",
	}
	data, _ := json.Marshal(cfg)
	reader := bytes.NewReader(data)

	req, err := http.NewRequest("POST", "/getAllProjects", reader)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleConfigSave())

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status code %v, but got %v", http.StatusOK, status)
	}
}
