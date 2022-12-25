package project

import (
	"encoding/json"

	"github.com/TimoSto/ThesorTeX/backend/app/internal/database"
)

func GetAllProjects(store database.ThesorTeXStore) ([]ProjectMetaDataWithName, error) {
	projectNames, err := store.GetAllProjectPaths()
	if err != nil {
		return nil, err
	}

	var projectsData []ProjectMetaDataWithName

	for _, p := range projectNames {
		file, err := store.ReadFileInProject(p, metaDataFile)
		if err != nil {
			return nil, err
		}
		var data ProjectMetaDataWithName
		err = json.Unmarshal(file, &data)
		if err != nil {
			return nil, err
		}

		data.Name = p

		projectsData = append(projectsData, data)
	}

	return projectsData, nil
}
