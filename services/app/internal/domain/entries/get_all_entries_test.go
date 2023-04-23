package entries

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/TimoSto/ThesorTeX/pkg/backend/filesystem/fake"
	"github.com/TimoSto/ThesorTeX/pkg/backend/pathbuilder"
	"github.com/TimoSto/ThesorTeX/services/app/internal/config"
)

func TestGetAllEntries(t *testing.T) {
	fs := fake.FileSystem{}

	cfg := config.Config{
		ProjectsDir: "projects/",
	}

	existing := []Entry{
		{
			Key:      "t1",
			Category: "c1",
			Fields: []string{
				"t",
				"aa",
				"2",
			},
		},
		{
			Key:      "t2",
			Category: "c2",
			Fields: []string{
				"t2",
				"aa2",
				"22",
			},
		},
	}
	data, _ := json.Marshal(existing)

	fs.WriteFile(pathbuilder.GetPathInProject(cfg.ProjectsDir, "test", entriesFile), data)

	got, err := GetAllEntries("test", &fs, cfg)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !reflect.DeepEqual(got, existing) {
		t.Errorf("expected %v, got %v", existing, got)
	}
}
