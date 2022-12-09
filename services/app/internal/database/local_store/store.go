package local_store

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"strings"

	"github.com/TimoSto/ThesorTeX/pkg/log"
	"github.com/TimoSto/ThesorTeX/pkg/pathbuilder"
	"github.com/TimoSto/ThesorTeX/services/app/conf"
	"github.com/TimoSto/ThesorTeX/services/app/internal/database"
)

type Store struct {
	Config conf.Config
}

const (
	projectDataFile = "project_data.json"
)

func (s *Store) GetAllProjects() ([]database.ProjectMetaData, error) {
	projectsPath := s.Config.ProjectsDir
	dirContent, err := os.ReadDir(projectsPath)
	if err != nil {
		err = os.Mkdir(projectsPath, 0644)
		if err != nil {
			return nil, err
		}
		dirContent, err = os.ReadDir(projectsPath)
		if err != nil {
			return nil, err
		}
	}

	var p []database.ProjectMetaData

	for _, f := range dirContent {
		if f.IsDir() {
			newProj := database.ProjectMetaData{
				Name: f.Name(),
			}
			confFile, err := ioutil.ReadFile(pathbuilder.GetPathInProject(s.Config.ProjectsDir, f.Name(), "config.json"))
			if err != nil {
				log.Error("got error while trying to read project config: %v", err)
				log.Info("Using empty config for project %s", f.Name())
			} else {
				var pConf database.ProjectMetaData
				err = json.Unmarshal(confFile, &pConf)
				if err != nil {
					log.Error("got error while trying to read project config: %v", err)
					log.Info("Using empty config for project %s", f.Name())
				} else {
					newProj.Created = pConf.Created
					newProj.LastModified = pConf.LastModified
				}
			}
			p = append(
				p,
				newProj,
			)
		}
	}

	return p, nil
}

func (s *Store) GetProjectData(project string) (database.ProjectData, error) {
	file, err := ioutil.ReadFile(pathbuilder.GetPathInProject(s.Config.ProjectsDir, project, projectDataFile))
	if err != nil {
		return database.ProjectData{}, err
	}

	var data database.ProjectData
	err = json.Unmarshal(file, &data)
	if err != nil {
		return database.ProjectData{}, err
	}

	return data, nil
}

func (s *Store) SaveProjectData(project string, data database.ProjectData) error {
	file, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(pathbuilder.GetPathInProject(s.Config.ProjectsDir, project, projectDataFile), file, 0777)
}

func (s *Store) CreateProject(metaData database.ProjectMetaData, template fs.FS) error {

	//todo:special error if project already exists
	err := os.Mkdir(pathbuilder.GetProjectPath(s.Config.ProjectsDir, metaData.Name), os.ModePerm)
	if err != nil {
		return err
	}

	err = os.Mkdir(pathbuilder.GetPathInProject(s.Config.ProjectsDir, metaData.Name, "styPackages"), os.ModePerm)
	if err != nil {
		return err
	}

	err = fs.WalkDir(template, ".", func(path string, d fs.DirEntry, err error) error {
		// cannot happen
		if err != nil {
			panic(err)
		}
		if d.IsDir() {
			return nil
		}
		b, err := fs.ReadFile(template, path)
		if err != nil {
			return err // or panic or ignore
		}
		path = strings.TrimPrefix(path, "template/")
		path = strings.Replace(path, "example.tex", fmt.Sprintf("%s.tex", metaData.Name), 1)

		return ioutil.WriteFile(pathbuilder.GetPathInProject(s.Config.ProjectsDir, metaData.Name, path), b, 0644)
	})

	if err != nil {
		return err
	}

	data, err := json.Marshal(metaData)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(pathbuilder.GetPathInProject(s.Config.ProjectsDir, metaData.Name, "config.json"), data, 0644)
}

func (s *Store) DeleteProject(name string) error {
	path := pathbuilder.GetProjectPath(s.Config.ProjectsDir, name)

	return os.RemoveAll(path)
}
