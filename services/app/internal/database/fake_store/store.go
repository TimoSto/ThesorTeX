package fake_store

import (
	"io/fs"

	"github.com/TimoSto/ThesorTeX/services/app/internal/database"
	"github.com/TimoSto/ThesorTeX/services/app/internal/project_template"
)

type Store struct {
	Entries     [][]database.BibEntry
	Categories  [][]database.BibCategory
	ProjectMeta []database.ProjectMetaData
}

var bibStyFile, _ = project_template.ProjectTemplate.ReadFile("template/styPackages/bibliography.sty")

func (s *Store) GetAllProjects() ([]database.ProjectMetaData, error) {
	return s.ProjectMeta, nil
}

func (s *Store) CreateProject(metaData database.ProjectMetaData, template fs.FS) error {
	s.ProjectMeta = append(s.ProjectMeta, metaData)

	return nil
}

func (s *Store) DeleteProject(name string) error {
	for i, p := range s.ProjectMeta {
		if p.Name == name {
			s.ProjectMeta = append(s.ProjectMeta[:i], s.ProjectMeta[i+1:]...)
			s.Entries = append(s.Entries[:i], s.Entries[i+1:]...)
			s.Categories = append(s.Categories[:i], s.Categories[i+1:]...)
			return nil
		}
	}

	return nil
}

func (s *Store) GetProjectEntries(project string) ([]database.BibEntry, error) {
	for i, p := range s.ProjectMeta {
		if p.Name == project {
			return s.Entries[i], nil
		}
	}

	return nil, nil
}

func (s *Store) SaveProjectEntries(project string, data []database.BibEntry) error {
	for i, p := range s.ProjectMeta {
		if p.Name == project {
			s.Entries[i] = data
			return nil
		}
	}

	return nil
}

func (s *Store) WriteCSV(project string, file []byte) error {
	return nil
}

func (s *Store) GetProjectCategories(project string) ([]database.BibCategory, error) {
	for i, p := range s.ProjectMeta {
		if p.Name == project {
			return s.Categories[i], nil
		}
	}

	return nil, nil
}

func (s *Store) SaveProjectCategories(project string, data []database.BibCategory) error {
	for i, p := range s.ProjectMeta {
		if p.Name == project {
			s.Categories[i] = data
			return nil
		}
	}

	return nil
}

func (s *Store) GetBibliographySty(project string) (string, error) {

	return string(bibStyFile), nil
}
func (s *Store) WriteBibliographySty(project string, file []byte) error {
	bibStyFile = file
	return nil
}
