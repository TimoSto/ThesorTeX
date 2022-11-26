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

func TestCreateProject(t *testing.T) {
	c := conf.Config{
		Port:        "",
		ProjectsDir: "/test2/",
	}
	p, err := CreateProject("testproject3", c, mock_fs.Mkdir, mock_fs.WriteFile)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}
	expectedFiles := []string{
		"/test2/testproject3/config.json",
		"/test2/testproject3/testproject3.tex",
		"/test2/testproject3/literatur.csv",
		"/test2/testproject3/literature_types.json",
		"/test2/testproject3/styPackages/abk_verzeichnis.sty",
		"/test2/testproject3/styPackages/anhang.sty",
		"/test2/testproject3/styPackages/codes.sty",
		"/test2/testproject3/styPackages/fusszeilen.sty",
		"/test2/testproject3/styPackages/header_footer.sty",
		"/test2/testproject3/styPackages/html.sty",
		"/test2/testproject3/styPackages/inhaltsverzeichnis.sty",
		"/test2/testproject3/styPackages/literatur.sty",
		"/test2/testproject3/styPackages/ueberschriften.sty",
	}
	for _, f := range expectedFiles {
		found := false
		for _, wf := range mock_fs.WrittenFiles {
			if wf == f {
				found = true
			}
		}
		if !found {
			t.Errorf("expected %s to have been written, but it wasnt", f)
		}
	}
	for _, f := range mock_fs.WrittenFiles {
		found := false
		for _, ef := range expectedFiles {
			if ef == f {
				found = true
			}
		}
		if !found {
			t.Errorf("unexpected file %s written", f)
		}
	}
	if p.Name != "testproject3" {
		t.Errorf("expected name, got %s", p.Name)
	}
}
