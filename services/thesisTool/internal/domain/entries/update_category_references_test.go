package entries

import (
	"encoding/json"
	"github.com/TimoSto/ThesorTeX/pkg/backend/filesystem/fake"
	"github.com/TimoSto/ThesorTeX/pkg/backend/pathbuilder"
	"github.com/TimoSto/ThesorTeX/services/thesisTool/internal/config"
	"testing"
)

func TestUpdateCategoryReferences(t *testing.T) {
	fs := fake.FileSystem{}

	cfg := config.Config{ProjectsDir: "/projects"}

	file := "[\n  {\n    \"Key\": \"testEntry\",\n    \"Category\": \"aufsatz\",\n    \"Fields\": [\n      \"Autor xyz\",\n      \"1999\",\n      \"ThesorTeX - Ein tolles Tool\",\n      \"Random Zeitschrift\",\n      \"99/03\",\n      \"20-22\",\n      \"https://wikipedia.org/\"\n    ]\n  }\n]"

	fs.WriteFile(pathbuilder.GetPathInProject(cfg.ProjectsDir, "test", entriesFile), []byte(file))

	err := UpdateCategoryReferences(&fs, cfg, "test", "aufsatz", "neuerAufsatz")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	resultFile, _ := fs.ReadFile(pathbuilder.GetPathInProject(cfg.ProjectsDir, "test", entriesFile))

	var result []Entry
	json.Unmarshal(resultFile, &result)

	if result[0].Category != "neuerAufsatz" {
		t.Errorf("expected entry to have category %s but got %s", "neuerAufsatz", result[0].Category)
	}
}
