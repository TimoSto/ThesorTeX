package entry

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TimoSto/ThesorTeX/services/app/internal/config"
	"github.com/TimoSto/ThesorTeX/services/app/internal/filesystem/fake"
)

func TestHandleUploadEntries_GET(t *testing.T) {
	req, err := http.NewRequest("GET", "/uploadEntries", nil)
	if err != nil {
		t.Fatal(err)
	}

	fs := fake.FileSystem{}

	config.Cfg = config.Config{
		ProjectsDir: "projects/",
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleUploadEntries(&fs))

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected %v, got %v", http.StatusMethodNotAllowed, rr.Code)
	}
}

func TestHandleUploadEntries_MissingQueryParameters(t *testing.T) {
	req, err := http.NewRequest("POST", "/saveCategory", nil)
	if err != nil {
		t.Fatal(err)
	}

	fs := fake.FileSystem{}

	config.Cfg = config.Config{
		ProjectsDir: "projects/",
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleUploadEntries(&fs))

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected %v, got %v", http.StatusBadRequest, rr.Code)
	}
}
