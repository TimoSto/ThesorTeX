package entries

import (
	"encoding/json"

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

	return nil
}
