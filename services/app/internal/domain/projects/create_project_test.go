package projects

import (
	"fmt"
	goFs "io/fs"
	"reflect"
	"strings"
	"testing"

	"github.com/TimoSto/ThesorTeX/services/app/internal/config"
	"github.com/TimoSto/ThesorTeX/services/app/internal/filesystem/fake"
	"github.com/TimoSto/ThesorTeX/services/app/internal/project_template"
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
			expectedRes: []string{"projects/test1", "projects/test1/data", "projects/test1/styPackages"},
			expectedErr: "",
		},
		{
			title:       "different existing",
			existing:    []string{"projects/test"},
			name:        "test1",
			expectedRes: []string{"projects/test", "projects/test1", "projects/test1/data", "projects/test1/styPackages"},
			expectedErr: "",
		},
		{
			title:       "same existing",
			existing:    []string{"projects/test1"},
			name:        "test1",
			expectedRes: []string{"projects/test1"},
			expectedErr: fmt.Sprintf(ErrorProjectPathAlreadyExists, "test1"),
		},
	}

	for _, s := range scenarios {
		t.Run(s.title, func(t *testing.T) {
			fs := fake.FileSystem{}
			for _, p := range s.existing {
				fs.CreateDirectory(p)
			}

			cfg := config.Config{ProjectsDir: "projects/"}

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
