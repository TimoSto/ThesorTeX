package projects

import (
	"reflect"
	"testing"

	"github.com/TimoSto/ThesorTeX/mocks/mock_fs"
	"github.com/TimoSto/ThesorTeX/services/app/conf"
)

func TestGetAllProjects_NoProjects(t *testing.T) {
	c := conf.Config{
		Port:        "",
		ProjectsDir: "/test0/",
	}
	projects, err := GetAllProjects(c, mock_fs.ReadDir, mock_fs.Mkdir, mock_fs.ReadFile)
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
	projects, err := GetAllProjects(c, mock_fs.ReadDir, mock_fs.Mkdir, mock_fs.ReadFile)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	if len(projects) != 2 {
		t.Errorf("Expected 2 projects, got %v", projects)
	}
	p1 := Project{
		Name:            "testproject1",
		Created:         "2022-11-13",
		LastModified:    "2022-11-20",
		NumberOfEntries: 0,
	}
	if !reflect.DeepEqual(p1, projects[0]) {
		t.Errorf("Expected %v, got %v", p1, projects[0])
	}
	p2 := Project{
		Name:            "testproject2",
		Created:         "2022-11-14",
		LastModified:    "2022-11-20",
		NumberOfEntries: 0,
	}
	if !reflect.DeepEqual(p2, projects[1]) {
		t.Errorf("Expected %v, got %v", p2, projects[1])
	}
}
