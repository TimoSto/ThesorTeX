package projects

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/TimoSto/ThesorTeX/services/app/internal/filesystem/fake"
)

func TestCreateProject(t *testing.T) {
	scenarios := []struct {
		title       string
		existing    []string
		name        string
		expectedRes []string
		expectedErr error
	}{
		{
			title:       "empty",
			existing:    []string{},
			name:        "test1",
			expectedRes: []string{"projects/test1"},
			expectedErr: nil,
		},
		{
			title:       "different existing",
			existing:    []string{"projects/test"},
			name:        "test1",
			expectedRes: []string{"projects/test", "projects/test1"},
			expectedErr: nil,
		},
		{
			title:       "same existing",
			existing:    []string{"projects/test1"},
			name:        "test1",
			expectedRes: []string{"projects/test1"},
			expectedErr: fmt.Errorf(ErrorProjectPathAlreadyExists, "test1"),
		},
	}

	for _, s := range scenarios {
		t.Run(s.title, func(t *testing.T) {
			fs := fake.FileSystem{}
			for _, p := range s.existing {
				fs.CreateDirectory(p)
			}
			err := CreateProject(s.name, &fs)
			if !reflect.DeepEqual(err, s.expectedErr) {
				t.Errorf("expected %v, got %v", s.expectedErr, err)
			}
			projects, _ := fs.GetAllDirectoriesUnder("")
			if !reflect.DeepEqual(projects, s.expectedRes) {
				t.Errorf("expected %v, got %v", s.expectedRes, projects)
			}
		})
	}
}
