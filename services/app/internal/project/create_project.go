package project

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"strings"
	"time"

	"github.com/TimoSto/ThesorTeX/services/app/internal/database"
	"github.com/TimoSto/ThesorTeX/services/app/internal/project_template"
)

var (
	ProjectAlreadyExistsError = fmt.Errorf("project already exists")
)

func CreateProject(name string, store database.ThesorTeXStore) (ProjectMetaDataWithName, error) {
	existing, err := store.GetAllProjectPaths()
	if err != nil {
		return ProjectMetaDataWithName{}, err
	}

	for _, e := range existing {
		if e == name {
			return ProjectMetaDataWithName{}, ProjectAlreadyExistsError
		}
	}

	err = store.MakeNewProjectPath(name)
	if err != nil {
		return ProjectMetaDataWithName{}, err
	}

	err = fs.WalkDir(project_template.ProjectTemplate, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		b, err := fs.ReadFile(project_template.ProjectTemplate, path)
		if err != nil {
			return err // or panic or ignore
		}
		path = strings.TrimPrefix(path, "template/")
		path = strings.Replace(path, "example.tex", "main.tex", 1)

		return store.WriteFileInProject(name, path, b)
	})

	if err != nil {
		return ProjectMetaDataWithName{}, err
	}

	meta := ProjectMetaDataWithName{
		ProjectMetaData: ProjectMetaData{
			Created:         time.Now().Format("2006-01-02 15:04"),
			LastModified:    time.Now().Format("2006-01-02 15:04"),
			NumberOfEntries: 1,
		},
		Name: name,
	}

	dataFile, err := json.Marshal(meta)
	if err != nil {
		return ProjectMetaDataWithName{}, err
	}

	err = store.WriteFileInProject(name, metaDataFile, dataFile)
	if err != nil {
		return ProjectMetaDataWithName{}, err
	}

	return meta, nil
}
