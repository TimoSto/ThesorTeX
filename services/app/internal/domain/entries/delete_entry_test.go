package entries

import (
	"encoding/json"
	"testing"

	"github.com/TimoSto/ThesorTeX/pkg/backend/pathbuilder"
	"github.com/TimoSto/ThesorTeX/services/app/internal/config"
	"github.com/TimoSto/ThesorTeX/services/app/internal/domain/projects"
	"github.com/TimoSto/ThesorTeX/services/app/internal/filesystem/fake"
)

func TestDeleteEntry(t *testing.T) {
	fs := fake.FileSystem{}

	cfg := config.Config{
		ProjectsDir: "/projects",
	}

	//TODO: generalize fake setup
	projects.CreateProject("test", &fs, cfg)

	err := DeleteEntry("test", "testEntry", &fs, cfg)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	file, _ := fs.ReadFile(pathbuilder.GetPathInProject(cfg.ProjectsDir, "test", entriesFile))
	var cs []Entry
	err = json.Unmarshal(file, &cs)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(cs) != 0 {
		t.Errorf("expected 0 entries but got %v", cs)
	}
}
