package project

import (
	"encoding/json"
	"time"

	"github.com/TimoSto/ThesorTeX/backend/app/internal/database"
)

func ReadMetaData(project string, store database.ThesorTeXStore) (ProjectMetaData, error) {
	dataFile, err := store.ReadFileInProject(project, metaDataFile)
	if err != nil {
		return ProjectMetaData{}, err
	}
	var data ProjectMetaData
	err = json.Unmarshal(dataFile, &data)
	if err != nil {
		return ProjectMetaData{}, err
	}
	return data, nil
}

func SaveMetaDataToJSON(project string, data ProjectMetaData, store database.ThesorTeXStore) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = store.WriteFileInProject(project, metaDataFile, jsonData)
	if err != nil {
		return err
	}

	return nil
}

func SaveProjectMetaData(project string, data ProjectMetaData, store database.ThesorTeXStore) error {
	data.LastModified = time.Now().Format("2006-01-02 15:04")

	err := SaveMetaDataToJSON(project, data, store)
	if err != nil {
		return err
	}

	return nil
}

func UpdateProjectLastEdited(project string, store database.ThesorTeXStore) (ProjectMetaData, error) {
	data, err := ReadMetaData(project, store)

	data.LastModified = time.Now().Format("2006-01-02 15:04")

	err = SaveProjectMetaData(project, data, store)
	if err != nil {
		return ProjectMetaData{}, err
	}

	return data, nil
}
