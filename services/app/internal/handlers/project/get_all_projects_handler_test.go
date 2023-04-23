package project

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/TimoSto/ThesorTeX/pkg/backend/filesystem/fake"
	"github.com/TimoSto/ThesorTeX/pkg/backend/pathbuilder"
	"github.com/TimoSto/ThesorTeX/services/app/internal/config"
	"github.com/TimoSto/ThesorTeX/services/app/internal/domain/projects"
)

func TestGetAllProjectsHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/getAllProjects", nil)
	if err != nil {
		t.Fatal(err)
	}

	fs := fake.FileSystem{}

	config.Cfg = config.Config{
		ProjectsDir: "projects",
	}

	existingData := projects.ProjectMetaData{
		Name:            "test1",
		Created:         time.Date(2000, 1, 1, 8, 0, 0, 0, time.Local).Format("2006-01-02 15:04"),
		LastEdited:      time.Date(2001, 1, 1, 8, 0, 0, 0, time.Local).Format("2006-01-02 15:04"),
		NumberOfEntries: 2,
	}

	data, _ := json.Marshal(existingData)
	path := pathbuilder.GetPathInProject(config.Cfg.ProjectsDir, "test", "data/metaData.json")

	fs.WriteFile(path, data)
	fs.CreateDirectory("projects/test")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAllProjectsHandler(&fs))

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status code %v, but got %v", http.StatusOK, status)
	}

	var p []projects.ProjectMetaData

	decoder := json.NewDecoder(rr.Body)
	err = decoder.Decode(&p)
	if err != nil {
		t.Fatal(err)
	}

	if len(p) != 1 {
		t.Errorf("expected one project data set, got %v", len(p))
	}
}
