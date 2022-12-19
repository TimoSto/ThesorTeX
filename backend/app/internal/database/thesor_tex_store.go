package database

import "io/fs"

type ThesorTeXStore interface {
	GetAllProjects() ([]ProjectMetaData, error)
	CreateProject(metaData ProjectMetaData, template fs.FS) error
	DeleteProject(project string) error
	GetProjectEntries(project string) ([]BibEntry, error)
	SaveProjectEntries(project string, data []BibEntry) error
	WriteCSV(project string, file []byte) error
	GetProjectCategories(project string) ([]BibCategory, error)
	SaveProjectCategories(project string, data []BibCategory) error
	GetBibliographySty(project string) (string, error)
	WriteBibliographySty(project string, file []byte) error
	SaveProjectMetaData(project string, data ProjectMetaData) error
	GetProjectMetaData(project string) (ProjectMetaData, error)
}
