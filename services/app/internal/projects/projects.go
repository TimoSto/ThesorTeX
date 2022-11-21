package projects

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"strings"
	"time"

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

//go:embed project_template
var projectTemplate embed.FS

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

func CreateProject(
	name string,
	config conf.Config,
	mkdir func(string, os.FileMode) error,
	writeFile func(string, []byte, fs.FileMode) error,
) error {

	//todo:special error if project already exists
	err := mkdir(pathbuilder.GetProjectPath(config.ProjectsDir, name), os.ModePerm)
	if err != nil {
		return err
	}

	err = fs.WalkDir(projectTemplate, ".", func(path string, d fs.DirEntry, err error) error {
		// cannot happen
		if err != nil {
			panic(err)
		}
		if d.IsDir() {
			return nil
		}
		b, err := fs.ReadFile(projectTemplate, path)
		if err != nil {
			return err // or panic or ignore
		}
		path = strings.TrimPrefix(path, "project_template/")
		path = strings.Replace(path, "example.tex", fmt.Sprintf("%s.tex", name), 1)

		return writeFile(pathbuilder.GetPathInProject(config.ProjectsDir, name, path), b, 0644)
	})

	if err != nil {
		return err
	}

	pObj := Project{
		Name:            "",
		Created:         time.Now().Format("2006-01-02 15:04"),
		LastModified:    time.Now().Format("2006-01-02 15:04"),
		NumberOfEntries: 0,
	}

	data, err := json.Marshal(pObj)
	if err != nil {
		return err
	}

	return writeFile(pathbuilder.GetPathInProject(config.ProjectsDir, name, "config.json"), data, 0644)
}
