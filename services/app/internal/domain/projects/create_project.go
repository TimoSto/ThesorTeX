package projects

import (
	"fmt"
	goFs "io/fs"
	"strings"

	"github.com/TimoSto/ThesorTeX/pkg/backend/pathbuilder"
	"github.com/TimoSto/ThesorTeX/services/app/internal/config"
	"github.com/TimoSto/ThesorTeX/services/app/internal/filesystem"
	"github.com/TimoSto/ThesorTeX/services/app/internal/project_template"
)

const (
	ErrorProjectPathAlreadyExists = "project path for %s already exists"
)

func CreateProject(name string, fs filesystem.FileSystem, cfg config.Config) error {
	path := pathbuilder.GetProjectPath(cfg.ProjectsDir, name)

	exists, err := fs.CheckDirectoryExists(path)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf(ErrorProjectPathAlreadyExists, name)
	}

	err = fs.CreateDirectory(path)
	if err != nil {
		return err
	}

	dataPath := pathbuilder.GetPathInProject(cfg.ProjectsDir, name, "/data")
	err = fs.CreateDirectory(dataPath)
	if err != nil {
		return err
	}

	styPath := pathbuilder.GetPathInProject(cfg.ProjectsDir, name, "/styPackages")
	err = fs.CreateDirectory(styPath)
	if err != nil {
		return err
	}

	err = goFs.WalkDir(project_template.ProjectTemplate, ".", func(path string, d goFs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		b, err := goFs.ReadFile(project_template.ProjectTemplate, path)
		if err != nil {
			return err // or panic or ignore
		}
		path = strings.TrimPrefix(path, "template/")

		filepath := pathbuilder.GetPathInProject(cfg.ProjectsDir, name, path)

		return fs.WriteFile(filepath, b)
	})

	return nil
}
