package categories

import (
	"encoding/json"
	"fmt"
	"github.com/TimoSto/ThesorTeX/services/thesisTool/internal/domain/projects"
	"sort"

	"github.com/TimoSto/ThesorTeX/pkg/backend/filesystem"
	"github.com/TimoSto/ThesorTeX/pkg/backend/pathbuilder"
	"github.com/TimoSto/ThesorTeX/services/thesisTool/internal/config"
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

	data, err := json.MarshalIndent(all, "", "\t")
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

	err = projects.UpdateProjectMetaData(fs, cfg, project, -1)
	if err != nil {
		return fmt.Errorf("got error updating the meta data: %v", err)
	}

	return nil
}
