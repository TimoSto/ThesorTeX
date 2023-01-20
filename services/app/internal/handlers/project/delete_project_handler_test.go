package project

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TimoSto/ThesorTeX/services/app/internal/config"
	"github.com/TimoSto/ThesorTeX/services/app/internal/filesystem/fake"
)

func TestHandleProjectDelete_GET(t *testing.T) {
	req, err := http.NewRequest("GET", "/createProject", nil)
	if err != nil {
		t.Fatal(err)
	}

	fs := fake.FileSystem{}

	cfg := config.Config{
		ProjectsDir: "projects/",
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleProjectDelete(&fs, cfg))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("expected status code %v, but got %v", http.StatusMethodNotAllowed, status)
	}
}

func TestHandleProjectDelete_MissingParameter(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/deleteProject", nil)
	if err != nil {
		t.Fatal(err)
	}

	fs := fake.FileSystem{}

	cfg := config.Config{
		ProjectsDir: "projects/",
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleProjectDelete(&fs, cfg))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("expected status code %v, but got %v", http.StatusBadRequest, status)
	}
}
