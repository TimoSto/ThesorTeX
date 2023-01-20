package projects

import (
	"github.com/TimoSto/ThesorTeX/pkg/backend/pathbuilder"
	"github.com/TimoSto/ThesorTeX/services/app/internal/config"
	"github.com/TimoSto/ThesorTeX/services/app/internal/filesystem"
)

func DeleteProject(name string, fs filesystem.FileSystem, cfg config.Config) error {
	path := pathbuilder.GetProjectPath(cfg.ProjectsDir, name)

	exists, err := fs.CheckDirectoryExists(path)
	if err != nil {
		return err
	}

	if exists {
		err = fs.DeleteDirectory(path)
		if err != nil {
			return err
		}
	}

	return nil
}
