package fake_store

import (
	"fmt"

	"github.com/TimoSto/ThesorTeX/services/app/internal/database"
)

type Store struct {
	Projects    []database.ProjectData
	ProjectMeta []database.ProjectMetaData
}

func (s *Store) GetProjectData(project string) (database.ProjectData, error) {
	for i, p := range s.ProjectMeta {
		if p.Name == project {
			return s.Projects[i], nil
		}
	}

	return database.ProjectData{}, fmt.Errorf("could not find project %v", project)
}

func (s *Store) SaveProjectData(project string, data database.ProjectData) error {
	for i, p := range s.ProjectMeta {
		if p.Name == project {
			s.Projects[i].Entries = data.Entries
			s.Projects[i].Categories = data.Categories
			return nil
		}
	}

	return nil
}
