package projects

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/TimoSto/ThesorTeX/pkg/backend/pathbuilder"
	"github.com/TimoSto/ThesorTeX/services/app/internal/config"
	"github.com/TimoSto/ThesorTeX/services/app/internal/filesystem/fake"
)

func TestGetAllProjects(t *testing.T) {
	fs := fake.FileSystem{}

	cfg := config.Config{
		ProjectsDir: "projects",
	}

	existingData := ProjectMetaData{
		Name:            "test1",
		Created:         time.Date(2000, 1, 1, 8, 0, 0, 0, time.Local).Format("2006-01-02 15:04"),
		LastEdited:      time.Date(2001, 1, 1, 8, 0, 0, 0, time.Local).Format("2006-01-02 15:04"),
		NumberOfEntries: 2,
	}

	data, _ := json.Marshal(existingData)
	path := pathbuilder.GetPathInProject(cfg.ProjectsDir, "test", metaDataFile)

	fs.WriteFile(path, data)
	fs.CreateDirectory("projects/test")

	projects, err := GetAllProjects(&fs, cfg)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(projects) != 1 {
		t.Errorf("expected to find 1 project, got %v", len(projects))
	}

	if !reflect.DeepEqual(projects[0], existingData) {
		t.Errorf("expected%v, got %v", existingData, projects[0])
	}
}
