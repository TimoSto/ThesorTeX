package entries

import (
	"encoding/json"
	"github.com/TimoSto/ThesorTeX/services/thesisTool/internal/domain/categories"
	"github.com/TimoSto/ThesorTeX/services/thesisTool/internal/domain/projects"

	"github.com/TimoSto/ThesorTeX/pkg/backend/filesystem"
	"github.com/TimoSto/ThesorTeX/pkg/backend/pathbuilder"
	"github.com/TimoSto/ThesorTeX/services/thesisTool/internal/config"
)

func DeleteEntry(project string, key string, fs filesystem.FileSystem, cfg config.Config) error {
	all, err := GetAllEntries(project, fs, cfg)
	if err != nil {
		return err
	}

	for i, c := range all {
		if c.Key == key {
			all = append(all[:i], all[i+1:]...)
		}
	}

	data, err := json.Marshal(all)
	if err != nil {
		return err
	}

	err = fs.WriteFile(pathbuilder.GetPathInProject(cfg.ProjectsDir, project, entriesFile), data)
	if err != nil {
		return err
	}

	avCategories, err := categories.GetAllCategories(project, fs, cfg)

	file := GenerateCsvForEntries(all, avCategories)

	err = fs.WriteFile(pathbuilder.GetPathInProject(cfg.ProjectsDir, project, csvFile), []byte(file))
	if err != nil {
		return err
	}

	err = projects.UpdateProjectMetaData(fs, cfg, project, len(all))
	if err != nil {
		return err
	}

	return nil
}
