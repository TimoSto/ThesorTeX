package categories

import (
	"encoding/json"
	"sort"

	"github.com/TimoSto/ThesorTeX/pkg/backend/pathbuilder"
	"github.com/TimoSto/ThesorTeX/services/app/internal/config"
	"github.com/TimoSto/ThesorTeX/services/app/internal/filesystem"
)

func SaveCategory(fs filesystem.FileSystem, cfg config.Config, project string, name string, category Category) error {
	all, err := GetAllCategories(project, fs, cfg)
	if err != nil {
		return err
	}

	found := false
	if len(name) > 0 {
		for i, c := range all {
			if c.Name == name {
				all[i] = category
				found = true
				break
			}
		}
	}

	if !found {
		all = append(all, category)
	}

	sort.SliceStable(all, func(i, j int) bool {
		return all[i].Name < all[j].Name
	})

	data, err := json.Marshal(all)
	if err != nil {
		return err
	}

	err = fs.WriteFile(pathbuilder.GetPathInProject(cfg.ProjectsDir, project, categoriesFile), data)
	if err != nil {
		return err
	}

	err = SaveCategoriesToSty(fs, cfg, project, all)
	if err != nil {
		return err
	}

	return nil
}
