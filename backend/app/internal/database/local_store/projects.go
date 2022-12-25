package local_store

import (
	"os"

	"github.com/TimoSto/ThesorTeX/backend/pkg/pathbuilder"
)

func (s *Store) GetAllProjectPaths() ([]string, error) {
	projectsPath := s.Config.ProjectsDir
	dirContent, err := os.ReadDir(projectsPath)
	if err != nil {
		err = os.Mkdir(projectsPath, 0700)
		if err != nil {
			return nil, err
		}
		dirContent, err = os.ReadDir(projectsPath)
		if err != nil {
			return nil, err
		}
	}

	var p []string

	for _, c := range dirContent {
		if c.IsDir() {
			p = append(p, c.Name())
		}
	}

	return p, nil
}

func (s *Store) MakeNewProjectPath(project string) error {
	//todo:special error if project already exists
	err := os.MkdirAll(pathbuilder.GetProjectPath(s.Config.ProjectsDir, project), 0700)
	if err != nil {
		return err
	}

	err = os.Mkdir(pathbuilder.GetPathInProject(s.Config.ProjectsDir, project, "styPackages"), 0700)
	if err != nil {
		return err
	}

	err = os.Mkdir(pathbuilder.GetPathInProject(s.Config.ProjectsDir, project, "data"), 0700)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) RemoveProject(project string) error {
	path := pathbuilder.GetProjectPath(s.Config.ProjectsDir, project)

	err := os.RemoveAll(path)

	if err != nil {
		return err
	}

	return nil
}
