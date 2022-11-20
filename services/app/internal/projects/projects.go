package projects

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"

	"github.com/TimoSto/ThesorTeX/pkg/log"
	"github.com/TimoSto/ThesorTeX/pkg/pathbuilder"
	"github.com/TimoSto/ThesorTeX/services/app/conf"
)

type Project struct {
	Name            string
	Created         string
	LastModified    string
	NumberOfEntries int
}

func GetAllProjects(
	config conf.Config,
	readDir func(string) ([]fs.FileInfo, error),
	mkdir func(string, os.FileMode) error,
	readFile func(string) ([]byte, error),
) ([]Project, error) {

	projectsPath := config.ProjectsDir
	dirContent, err := readDir(projectsPath)
	if err != nil {
		err = mkdir(projectsPath, 0644)
		if err != nil {
			return nil, err
		}
		dirContent, err = readDir(projectsPath)
		if err != nil {
			return nil, err
		}
	}

	var projects []Project

	for _, f := range dirContent {
		if f.IsDir() {
			newProj := Project{
				Name: f.Name(),
			}
			confFile, err := readFile(pathbuilder.GetPathInProject(config.ProjectsDir, f.Name(), "config.json"))
			if err != nil {
				log.Error(fmt.Sprintf("got error while trying to read project config: %v", err))
				log.Info(fmt.Sprintf("Using empty config for project %s", f.Name()))
			} else {
				var pConf Project
				err = json.Unmarshal(confFile, &pConf)
				if err != nil {
					log.Error(fmt.Sprintf("got error while trying to read project config: %v", err))
					log.Info(fmt.Sprintf("Using empty config for project %s", f.Name()))
				} else {
					newProj.Created = pConf.Created
					newProj.LastModified = pConf.LastModified
				}
			}
			projects = append(
				projects,
				newProj,
			)
		}
	}

	return projects, nil
}
