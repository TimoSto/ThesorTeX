package categories

import (
	"encoding/json"

	"github.com/TimoSto/ThesorTeX/pkg/backend/filesystem"
	"github.com/TimoSto/ThesorTeX/pkg/backend/pathbuilder"
	"github.com/TimoSto/ThesorTeX/services/thesisTool/internal/config"
)

func GetAllCategories(project string, fs filesystem.FileSystem, cfg config.Config) ([]Category, error) {
	file, err := fs.ReadFile(pathbuilder.GetPathInProject(cfg.ProjectsDir, project, categoriesFile))
	if err != nil {
		return nil, err
	}

	var categories []Category
	err = json.Unmarshal(file, &categories)
	if err != nil {
		return nil, err
	}

	return categories, nil
}
