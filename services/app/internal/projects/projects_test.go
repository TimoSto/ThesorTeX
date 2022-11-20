package projects

import (
	"testing"

	"github.com/TimoSto/ThesorTeX/mocks/mock_fs"
	"github.com/TimoSto/ThesorTeX/services/app/conf"
)

func TestGetAllProjects_NoProjects(t *testing.T) {
	c := conf.Config{
		Port:        "",
		ProjectsDir: "/test0/",
	}
	projects, err := GetAllProjects(c, mock_fs.ReadDir, mock_fs.Mkdir)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if len(projects) != 0 {
		t.Errorf("Expected [], got %v", projects)
	}
}

func TestGetAllProjects_TwoProjects(t *testing.T) {
	c := conf.Config{
		Port:        "",
		ProjectsDir: "/test2/",
	}
	projects, err := GetAllProjects(c, mock_fs.ReadDir, mock_fs.Mkdir)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if len(projects) != 2 {
		t.Errorf("Expected 2 projects, got %v", projects)
	}
}
