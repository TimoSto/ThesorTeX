package project

import (
	"time"

	"github.com/TimoSto/ThesorTeX/backend/app/internal/database"
)

func SaveProjectMetaData(project string, data database.ProjectMetaData, store database.ThesorTeXStore) error {
	data.LastModified = time.Now().Format("2006-01-02 15:04")

	err := store.SaveProjectMetaData(project, data)
	if err != nil {
		return err
	}

	return nil
}
