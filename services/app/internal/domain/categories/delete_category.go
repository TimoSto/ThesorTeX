package categories

import (
	"encoding/json"

	"github.com/TimoSto/ThesorTeX/pkg/backend/pathbuilder"
	"github.com/TimoSto/ThesorTeX/services/app/internal/config"
	"github.com/TimoSto/ThesorTeX/services/app/internal/filesystem"
)

func DeleteCategory(project string, name string, fs filesystem.FileSystem, cfg config.Config) error {
	all, err := GetAllCategories(project, fs, cfg)
	if err != nil {
		return err
	}

	for i, c := range all {
		if c.Name == name {
			all = append(all[:i], all[i+1:]...)
		}
	}

	data, err := json.Marshal(all)
	if err != nil {
		return err
	}

	err = fs.WriteFile(pathbuilder.GetPathInProject(cfg.ProjectsDir, project, categoriesFile), data)
	if err != nil {
		return err
	}

	return nil
}
