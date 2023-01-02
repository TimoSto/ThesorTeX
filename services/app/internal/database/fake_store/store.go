package fake_store

import (
	"github.com/TimoSto/ThesorTeX/pkg/pathbuilder"
	"github.com/TimoSto/ThesorTeX/services/app/internal/project_template"
)

type Store struct {
	Files    map[string][]byte
	Projects []string
}

var bibStyFile, _ = project_template.ProjectTemplate.ReadFile("template/styPackages/bibliography.sty")

func (s *Store) GetAllProjectPaths() ([]string, error) {
	return s.Projects, nil
}

func (s *Store) MakeNewProjectPath(project string) error {
	s.Projects = append(s.Projects, project)

	return nil
}

func (s *Store) RemoveProject(name string) error {
	for i, p := range s.Projects {
		if p == name {
			s.Projects = append(s.Projects[:i], s.Projects[i+1:]...)
			return nil
		}
	}

	return nil
}

func (s *Store) ReadFileInProject(project string, path string) ([]byte, error) {
	fullPath := pathbuilder.GetPathInProject("", project, path)

	file := s.Files[fullPath]

	return file, nil
}

func (s *Store) WriteFileInProject(project string, path string, file []byte) error {
	if s.Files == nil {
		s.Files = make(map[string][]byte)
	}
	fullPath := pathbuilder.GetPathInProject("", project, path)

	s.Files[fullPath] = file

	return nil
}
