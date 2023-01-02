package database

/**
In current state, this package is more of a filesystem handler/proxy
*/

type ThesorTeXStore interface {
	GetAllProjectPaths() ([]string, error)
	MakeNewProjectPath(project string) error
	RemoveProject(project string) error
	ReadFileInProject(project string, path string) ([]byte, error)
	WriteFileInProject(project string, path string, file []byte) error
}
