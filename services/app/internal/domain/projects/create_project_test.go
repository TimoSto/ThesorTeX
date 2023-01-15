package projects

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/TimoSto/ThesorTeX/services/app/internal/config"
	"github.com/TimoSto/ThesorTeX/services/app/internal/filesystem/fake"
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

			err := CreateProject(s.name, &fs, cfg)
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
		})
	}
}
