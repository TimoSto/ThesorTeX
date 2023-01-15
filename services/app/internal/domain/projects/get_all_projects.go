package projects

import (
	"encoding/json"

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
		file, err := fs.ReadFile(path)
		if err != nil {
			return nil, err
		}

		var data ProjectMetaData
		err = json.Unmarshal(file, &data)
		if err != nil {
			return nil, err
		}

		projects = append(projects, data)
	}

	return projects, nil
}
