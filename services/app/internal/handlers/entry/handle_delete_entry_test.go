package entry

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TimoSto/ThesorTeX/pkg/backend/filesystem/fake"
	"github.com/TimoSto/ThesorTeX/services/app/internal/config"
)

func TestHandleDeleteEntry_GET(t *testing.T) {
	req, err := http.NewRequest("GET", "/deleteEntry", nil)
	if err != nil {
		t.Fatal(err)
	}

	fs := fake.FileSystem{}

	config.Cfg = config.Config{
		ProjectsDir: "projects/",
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleDeleteEntry(&fs))

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected %v, got %v", http.StatusMethodNotAllowed, rr.Code)
	}
}

func TestHandleDeleteEntry_MissingQueryParameters(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/deleteEntry", nil)
	if err != nil {
		t.Fatal(err)
	}

	fs := fake.FileSystem{}

	config.Cfg = config.Config{
		ProjectsDir: "projects/",
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleDeleteEntry(&fs))

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected %v, got %v", http.StatusBadRequest, rr.Code)
	}
}
