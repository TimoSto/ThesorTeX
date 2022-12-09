package project

import (
	"fmt"
	"time"

	"github.com/TimoSto/ThesorTeX/services/app/internal/database"
	"github.com/TimoSto/ThesorTeX/services/app/internal/project_template"
)

var (
	ProjectAlreadyExistsError = fmt.Errorf("project already exists")
)

func CreateProject(name string, store database.ThesorTeXStore) (database.ProjectMetaData, error) {
	existing, err := store.GetAllProjects()
	if err != nil {
		return database.ProjectMetaData{}, err
	}

	for _, e := range existing {
		if e.Name == name {
			return database.ProjectMetaData{}, ProjectAlreadyExistsError
		}
	}

	meta := database.ProjectMetaData{
		Name:            name,
		Created:         time.Now().Format("2006-01-02 15:04"),
		LastModified:    time.Now().Format("2006-01-02 15:04"),
		NumberOfEntries: 0,
	}

	return meta, store.CreateProject(meta, project_template.ProjectTemplate)
}
