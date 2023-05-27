package entries

import (
	"encoding/json"
	"github.com/TimoSto/ThesorTeX/pkg/backend/filesystem"
	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"github.com/TimoSto/ThesorTeX/pkg/backend/pathbuilder"
	"github.com/TimoSto/ThesorTeX/services/thesisTool/internal/config"
	"github.com/TimoSto/ThesorTeX/services/thesisTool/internal/domain/categories"
)

func UpdateCategoryReferences(fs filesystem.FileSystem, cfg config.Config, project string, category string, newCategory string) error {
	if category == newCategory {
		// nothing to update
		return nil
	}

	log.Info("Updating references to category %s (previously %s) in project %s...", newCategory, category, project)

	all, err := GetAllEntries(project, fs, cfg)
	if err != nil {
		return err
	}

	for i, e := range all {
		if e.Category == category {
			all[i].Category = newCategory
		}
	}

	// TODO: generalize into func and then make consts private again
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

	return nil
}
