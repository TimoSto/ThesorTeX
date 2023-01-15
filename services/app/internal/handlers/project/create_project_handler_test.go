package project

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/TimoSto/ThesorTeX/services/app/internal/config"
	"github.com/TimoSto/ThesorTeX/services/app/internal/filesystem/fake"
)

func TestCreateProjectHandler_GET(t *testing.T) {
	req, err := http.NewRequest("GET", "/createProject", nil)
	if err != nil {
		t.Fatal(err)
	}

	fs := fake.FileSystem{}

	cfg := config.Config{
		ProjectsDir: "projects/",
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateProjectHandler(&fs, cfg))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("expected status code %v, but got %v", http.StatusBadRequest, status)
	}
}

func TestCreateProjectHandler_PUT(t *testing.T) {
	data := createProjectData{Name: "test"}
	body, _ := json.Marshal(data)
	reader := bytes.NewReader(body)
	req, err := http.NewRequest("PUT", "/createProject", reader)
	if err != nil {
		t.Fatal(err)
	}

	fs := fake.FileSystem{}

	cfg := config.Config{
		ProjectsDir: "projects/",
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateProjectHandler(&fs, cfg))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status code %v, but got %v", http.StatusOK, status)
	}

	existing, _ := fs.GetAllDirectoriesUnder("")
	if expected := []string{"projects/test", "projects/test/data", "projects/test/styPackages"}; !reflect.DeepEqual(existing, expected) {
		t.Errorf("expected %v but got %v", expected, existing)
	}
}
