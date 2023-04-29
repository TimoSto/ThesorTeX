package projects

import (
	"fmt"
	goFs "io/fs"
	"reflect"
	"strings"
	"testing"

	"github.com/TimoSto/ThesorTeX/pkg/backend/filesystem/fake"
	"github.com/TimoSto/ThesorTeX/pkg/backend/pathbuilder"
	"github.com/TimoSto/ThesorTeX/pkg/backend/project_template"
	"github.com/TimoSto/ThesorTeX/services/thesisTool/internal/config"
)

func TestCreateProject(t *testing.T) {
	scenarios := []struct {
		title       string
		existing    []string
		name        string
		expectedRes []string
		expectedErr string
	}{
		{
			title:       "empty",
			existing:    []string{},
			name:        "test1",
			expectedRes: []string{"test1"},
			expectedErr: "",
		},
		{
			title:       "different existing",
			existing:    []string{"test"},
			name:        "test1",
			expectedRes: []string{"test", "test1"},
			expectedErr: "",
		},
		{
			title:       "same existing",
			existing:    []string{"test1"},
			name:        "test1",
			expectedRes: []string{"test1"},
			expectedErr: fmt.Sprintf(ErrorProjectPathAlreadyExists, "test1"),
		},
	}

	for _, s := range scenarios {
		t.Run(s.title, func(t *testing.T) {
			fs := fake.FileSystem{}
			cfg := config.Config{ProjectsDir: "projects/"}

			for _, p := range s.existing {
				fs.CreateDirectory(pathbuilder.GetProjectPath(cfg.ProjectsDir, p))
			}

			_, err := CreateProject(s.name, &fs, cfg)
			if err == nil && s.expectedErr != "" {
				t.Errorf("expected %v, got %v", s.expectedErr, err)
			}

			errStr := ""
			if err != nil {
				errStr = err.Error()
			}

			if errStr != s.expectedErr {
				t.Errorf("expected %v, got %v", s.expectedErr, errStr)
			}
			projects, _ := fs.GetAllDirectoriesUnder("")
			if !reflect.DeepEqual(projects, s.expectedRes) {
				t.Errorf("expected %v, got %v", s.expectedRes, projects)
			}

			//check if all files from template are there
			if s.expectedErr == "" {
				goFs.WalkDir(project_template.ProjectTemplate, ".", func(path string, d goFs.DirEntry, err error) error {
					if !d.IsDir() {
						path = strings.TrimPrefix(path, "template/")
						path = fmt.Sprintf("projects/%s/%s", s.name, path)
						file, err := fs.ReadFile(path)
						if err != nil {
							t.Errorf("unexpected error %v", err)
						}
						if len(file) == 0 {
							t.Errorf("file %v is empty", path)
						}
					}

					return nil
				})
			}
		})
	}
}
