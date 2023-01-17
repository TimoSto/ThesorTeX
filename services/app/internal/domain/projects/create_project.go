package projects

import (
	"encoding/json"
	"fmt"
	goFs "io/fs"
	"strings"
	"time"

	"github.com/TimoSto/ThesorTeX/pkg/backend/pathbuilder"
	"github.com/TimoSto/ThesorTeX/services/app/internal/config"
	"github.com/TimoSto/ThesorTeX/services/app/internal/filesystem"
	"github.com/TimoSto/ThesorTeX/services/app/internal/project_template"
)

const (
	ErrorProjectPathAlreadyExists = "project path for %s already exists"

	metaDataFile = "/data/metaData.json"
)

func CreateProject(name string, fs filesystem.FileSystem, cfg config.Config) (ProjectMetaData, error) {
	path := pathbuilder.GetProjectPath(cfg.ProjectsDir, name)

	exists, err := fs.CheckDirectoryExists(path)
	if err != nil {
		return ProjectMetaData{}, err
	}
	if exists {
		return ProjectMetaData{}, fmt.Errorf(ErrorProjectPathAlreadyExists, name)
	}

	err = fs.CreateDirectory(path)
	if err != nil {
		return ProjectMetaData{}, err
	}

	dataPath := pathbuilder.GetPathInProject(cfg.ProjectsDir, name, "/data")
	err = fs.CreateDirectory(dataPath)
	if err != nil {
		return ProjectMetaData{}, err
	}

	styPath := pathbuilder.GetPathInProject(cfg.ProjectsDir, name, "/styPackages")
	err = fs.CreateDirectory(styPath)
	if err != nil {
		return ProjectMetaData{}, err
	}

	err = goFs.WalkDir(project_template.ProjectTemplate, ".", func(path string, d goFs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		b, err := goFs.ReadFile(project_template.ProjectTemplate, path)
		if err != nil {
			return err // or panic or ignore
		}
		path = strings.TrimPrefix(path, "template/")

		// todo: update time and also check that in tests

		filepath := pathbuilder.GetPathInProject(cfg.ProjectsDir, name, path)

		return fs.WriteFile(filepath, b)
	})

	meta := ProjectMetaData{
		Created:         time.Now().Format("2006-01-02 15:04"),
		LastEdited:      time.Now().Format("2006-01-02 15:04"),
		NumberOfEntries: 1,
		Name:            name,
	}

	dataFile, err := json.Marshal(meta)

	metaDataPath := pathbuilder.GetPathInProject(cfg.ProjectsDir, name, metaDataFile)

	err = fs.WriteFile(metaDataPath, dataFile)
	if err != nil {
		return ProjectMetaData{}, err
	}

	return meta, nil
}
