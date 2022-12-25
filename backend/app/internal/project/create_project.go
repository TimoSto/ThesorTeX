package project

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"strings"
	"time"

	"github.com/TimoSto/ThesorTeX/backend/app/internal/database"
	"github.com/TimoSto/ThesorTeX/backend/app/internal/project_template"
)

var (
	ProjectAlreadyExistsError = fmt.Errorf("project already exists")
)

func CreateProject(name string, store database.ThesorTeXStore) (ProjectMetaData, error) {
	existing, err := store.GetAllProjectPaths()
	if err != nil {
		return ProjectMetaData{}, err
	}

	for _, e := range existing {
		if e == name {
			return ProjectMetaData{}, ProjectAlreadyExistsError
		}
	}

	err = store.MakeNewProjectPath(name)
	if err != nil {
		return ProjectMetaData{}, err
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
		return ProjectMetaData{}, err
	}

	meta := ProjectMetaData{
		Created:         time.Now().Format("2006-01-02 15:04"),
		LastModified:    time.Now().Format("2006-01-02 15:04"),
		NumberOfEntries: 1,
	}

	dataFile, err := json.Marshal(meta)
	if err != nil {
		return ProjectMetaData{}, err
	}

	err = store.WriteFileInProject(name, metaDataFile, dataFile)
	if err != nil {
		return ProjectMetaData{}, err
	}
	
	return meta, nil
}
