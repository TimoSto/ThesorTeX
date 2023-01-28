package fake

type FileSystem struct {
	dirs  []string
	files map[string][]byte
}

func (fs *FileSystem) CheckDirectoryExists(path string) (bool, error) {
	for _, d := range fs.dirs {
		if d == path {
			return true, nil
		}
	}
	return false, nil
}

func (fs *FileSystem) CreateDirectory(path string) error {
	fs.dirs = append(fs.dirs, path)
	return nil
}

func (fs *FileSystem) DeleteDirectory(path string) error {
	for i, p := range fs.dirs {
		if p == path {
			fs.dirs = append(fs.dirs[:i], fs.dirs[i+1:]...)
			return nil
		}
	}
	return nil
}

func (fs *FileSystem) WriteFile(path string, content []byte) error {
	if fs.files == nil {
		fs.files = make(map[string][]byte)
	}

	fs.files[path] = content

	return nil
}

func (fs *FileSystem) ReadFile(path string) ([]byte, error) {

	return fs.files[path], nil
}

func (fs *FileSystem) GetAllDirectoriesUnder(path string) ([]string, error) {
	return fs.dirs, nil
}
