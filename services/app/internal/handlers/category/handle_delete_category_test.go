package category

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TimoSto/ThesorTeX/services/app/internal/config"
	"github.com/TimoSto/ThesorTeX/services/app/internal/filesystem/fake"
)

func TestHandleDeleteCategory_GET(t *testing.T) {
	req, err := http.NewRequest("GET", "/deleteCategory", nil)
	if err != nil {
		t.Fatal(err)
	}

	fs := fake.FileSystem{}

	cfg := config.Config{
		ProjectsDir: "projects/",
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleDeleteCategory(&fs, cfg))

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected %v, got %v", http.StatusMethodNotAllowed, rr.Code)
	}
}

func TestHandleDeleteCategory_MissingQueryParameters(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/deleteCategory", nil)
	if err != nil {
		t.Fatal(err)
	}

	fs := fake.FileSystem{}

	cfg := config.Config{
		ProjectsDir: "projects/",
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleDeleteCategory(&fs, cfg))

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected %v, got %v", http.StatusBadRequest, rr.Code)
	}
}