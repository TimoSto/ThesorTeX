package projects

import (
	"io/ioutil"
	"os"

	"github.com/TimoSto/ThesorTeX/services/app/conf"
)

type Project struct {
	Name string
}

func GetAllProjects() ([]Project, error) {
	projectsPath := conf.GetConfig().ProjectsDir
	dirContent, err := ioutil.ReadDir(projectsPath)
	if err != nil {
		err = os.Mkdir(projectsPath, 0644)
		if err != nil {
			return nil, err
		}
		dirContent, err = ioutil.ReadDir(projectsPath)
		if err != nil {
			return nil, err
		}
	}

	var projects []Project

	for _, f := range dirContent {
		if f.IsDir() {
			projects = append(
				projects,
				Project{
					Name: f.Name(),
				})
		}
	}

	return projects, nil
}