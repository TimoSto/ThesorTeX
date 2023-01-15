package local

import (
	"io/ioutil"
	"os"
)

type FileSystem struct{}

func (fs *FileSystem) CheckDirectoryExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func (fs *FileSystem) CreateDirectory(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func (fs *FileSystem) DeleteDirectory(path string) error {
	return os.RemoveAll(path)
}

func (fs *FileSystem) WriteFile(path string, content []byte) error {
	return os.WriteFile(path, content, os.ModePerm)
}

func (fs *FileSystem) ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func (fs *FileSystem) GetAllDirectoriesUnder(path string) ([]string, error) {
	res, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var dirs []string

	for _, r := range res {
		if r.IsDir() {
			dirs = append(dirs, r.Name())
		}
	}

	return dirs, nil
}
