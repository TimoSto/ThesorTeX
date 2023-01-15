package projects

import (
	"fmt"

	"github.com/TimoSto/ThesorTeX/pkg/backend/pathbuilder"
	"github.com/TimoSto/ThesorTeX/services/app/internal/filesystem"
)

const (
	ErrorProjectPathAlreadyExists = "project path for %s already exists"
)

func CreateProject(name string, fs filesystem.FileSystem) error {
	path := pathbuilder.GetProjectPath(pathbuilder.GetPathFromExecRoot("projects"), name)
	
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

	return nil
}
