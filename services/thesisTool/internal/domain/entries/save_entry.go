package entries

import (
	"encoding/json"
	"sort"

	"github.com/TimoSto/ThesorTeX/pkg/backend/filesystem"
	"github.com/TimoSto/ThesorTeX/pkg/backend/pathbuilder"
	"github.com/TimoSto/ThesorTeX/services/thesisTool/internal/config"
	"github.com/TimoSto/ThesorTeX/services/thesisTool/internal/domain/categories"
)

func SaveEntry(fs filesystem.FileSystem, cfg config.Config, project string, key string, entry Entry) error {
	all, err := GetAllEntries(project, fs, cfg)
	if err != nil {
		return err
	}

	found := false
	if len(key) > 0 {
		for i, c := range all {
			if c.Key == key {
				all[i] = entry
				found = true
				break
			}
		}
	}

	if !found {
		all = append(all, entry)
	}

	sort.SliceStable(all, func(i, j int) bool {
		return all[i].Key < all[j].Key
	})

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

func SaveEntries(fs filesystem.FileSystem, cfg config.Config, project string, entries []Entry) error {
	all, err := GetAllEntries(project, fs, cfg)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		found := false
		if len(entry.Key) > 0 {
			for i, c := range all {
				if c.Key == entry.Key {
					all[i] = entry
					found = true
					break
				}
			}
		}

		if !found {
			all = append(all, entry)
		}
	}

	sort.SliceStable(all, func(i, j int) bool {
		return all[i].Key < all[j].Key
	})

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
