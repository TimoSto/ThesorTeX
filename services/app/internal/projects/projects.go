package projects

import (
	"io/fs"
	"os"

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
			projects = append(
				projects,
				newProj,
			)
		}
	}

	return projects, nil
}
