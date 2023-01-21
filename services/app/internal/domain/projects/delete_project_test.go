package projects

import (
	"testing"

	"github.com/TimoSto/ThesorTeX/services/app/internal/config"
	"github.com/TimoSto/ThesorTeX/services/app/internal/filesystem/fake"
)

func TestDeleteProject(t *testing.T) {
	fs := fake.FileSystem{}
	cfg := config.Config{ProjectsDir: "/projects"}

	fs.CreateDirectory("/projects/test")

	err := DeleteProject("test", &fs, cfg)
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}

	if e, _ := fs.CheckDirectoryExists("/projects/test"); e {
		t.Errorf("expected not to find directory '/projects/test', but found it")
	}
}