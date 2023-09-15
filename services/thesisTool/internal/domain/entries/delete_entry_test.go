package entries

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/TimoSto/ThesorTeX/pkg/backend/filesystem/fake"
	"github.com/TimoSto/ThesorTeX/pkg/backend/pathbuilder"
	"github.com/TimoSto/ThesorTeX/services/thesisTool/internal/config"
	"github.com/TimoSto/ThesorTeX/services/thesisTool/internal/domain/projects"
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

	csvFile, _ := fs.ReadFile(pathbuilder.GetPathInProject(cfg.ProjectsDir, "test", csvFile))
	lines := strings.Split(strings.Trim(string(csvFile), "\n"), "\n")
	if len(lines) != 2 {
		t.Errorf("expected 2 lines, got: %v", lines)
	}
}
