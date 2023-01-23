package categories

import (
	"encoding/json"
	"testing"

	"github.com/TimoSto/ThesorTeX/pkg/backend/pathbuilder"
	"github.com/TimoSto/ThesorTeX/services/app/internal/config"
	"github.com/TimoSto/ThesorTeX/services/app/internal/domain/projects"
	"github.com/TimoSto/ThesorTeX/services/app/internal/filesystem/fake"
)

func TestDeleteCategory(t *testing.T) {
	fs := fake.FileSystem{}

	cfg := config.Config{
		ProjectsDir: "/projects",
	}

	//TODO: generalize fake setup
	projects.CreateProject("test", &fs, cfg)

	err := DeleteCategory("test", "aufsatz", &fs, cfg)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	file, _ := fs.ReadFile(pathbuilder.GetPathInProject(cfg.ProjectsDir, "test", categoriesFile))
	var cs []Category
	err = json.Unmarshal(file, &cs)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(cs) != 0 {
		t.Errorf("expected 0 categories but got %v", cs)
	}
}
