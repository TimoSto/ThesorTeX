package database

import "io/fs"

type ThesorTeXStore interface {
	GetAllProjects() ([]ProjectMetaData, error)
	GetProjectData(project string) (ProjectData, error)
	SaveProjectData(project string, data ProjectData) error
	CreateProject(metaData ProjectMetaData, template fs.FS) error
	DeleteProject(project string) error
}
