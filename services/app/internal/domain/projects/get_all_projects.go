package projects

import (
	"github.com/TimoSto/ThesorTeX/pkg/backend/pathbuilder"
	"github.com/TimoSto/ThesorTeX/services/app/internal/config"
	"github.com/TimoSto/ThesorTeX/services/app/internal/filesystem"
)

func GetAllProjects(fs filesystem.FileSystem, cfg config.Config) ([]ProjectMetaData, error) {
	dirs, err := fs.GetAllDirectoriesUnder(cfg.ProjectsDir)
	if err != nil {
		return nil, err
	}

	var projects []ProjectMetaData

	for _, d := range dirs {
		path := pathbuilder.GetPathInProject(cfg.ProjectsDir, d, metaDataFile)
		fs.ReadFile(path)
	}

	return projects, nil
}
