package entries

import (
	"encoding/json"

	"github.com/TimoSto/ThesorTeX/pkg/backend/filesystem"
	"github.com/TimoSto/ThesorTeX/pkg/backend/pathbuilder"
	"github.com/TimoSto/ThesorTeX/services/app/internal/config"
)

func GetAllEntries(project string, fs filesystem.FileSystem, cfg config.Config) ([]Entry, error) {
	file, err := fs.ReadFile(pathbuilder.GetPathInProject(cfg.ProjectsDir, project, entriesFile))
	if err != nil {
		return nil, err
	}

	var entries []Entry
	err = json.Unmarshal(file, &entries)
	if err != nil {
		return nil, err
	}

	return entries, nil
}
