package local_store

import (
	"os"

	"github.com/TimoSto/ThesorTeX/pkg/pathbuilder"
)

func (s *Store) ReadFileInProject(project string, path string) ([]byte, error) {
	fullPath := pathbuilder.GetPathInProject(s.Config.ProjectsDir, project, path)

	file, err := os.ReadFile(fullPath)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (s *Store) WriteFileInProject(project string, path string, file []byte) error {
	fullPath := pathbuilder.GetPathInProject(s.Config.ProjectsDir, project, path)

	err := os.WriteFile(fullPath, file, 0700)

	if err != nil {
		return err
	}

	return nil
}
