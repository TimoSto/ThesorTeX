package database

import "io/fs"

type ThesorTeXStore interface {
	GetAllProjects() ([]ProjectMetaData, error)
	CreateProject(metaData ProjectMetaData, template fs.FS) error
	DeleteProject(project string) error
	GetProjectEntries(project string) ([]BibEntry, error)
	SaveProjectEntries(project string, data []BibEntry) error
	GetProjectCategories(project string) ([]BibCategory, error)
	SaveProjectCategories(project string, data []BibCategory) error
}
