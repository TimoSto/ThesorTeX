package projects

import (
	"encoding/json"
	"github.com/TimoSto/ThesorTeX/pkg/backend/filesystem"
	"github.com/TimoSto/ThesorTeX/pkg/backend/pathbuilder"
	"github.com/TimoSto/ThesorTeX/services/thesisTool/internal/config"
	"time"
)

func UpdateProjectMetaData(fs filesystem.FileSystem, cfg config.Config, project string, entries int) error {
	metaDataPath := pathbuilder.GetPathInProject(cfg.ProjectsDir, project, metaDataFile)
	file, err := fs.ReadFile(metaDataPath)
	if err != nil {
		return err
	}

	var metaData ProjectMetaData
	err = json.Unmarshal(file, &metaData)
	if err != nil {
		return err
	}

	metaData.LastEdited = time.Now().Format("2006-01-02 15:04")
	if entries > -1 {
		metaData.NumberOfEntries = entries
	}

	data, err := json.Marshal(metaData)
	if err != nil {
		return err
	}

	err = fs.WriteFile(metaDataPath, data)
	if err != nil {
		return err
	}

	return nil
}
