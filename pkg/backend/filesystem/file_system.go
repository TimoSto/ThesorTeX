package filesystem

type FileSystem interface {
	CheckDirectoryExists(path string) (bool, error)
	GetAllDirectoriesUnder(path string) ([]string, error)
	CreateDirectory(path string) error
	DeleteDirectory(path string) error
	WriteFile(path string, content []byte) error
	ReadFile(path string) ([]byte, error)
}
