package mock_fs

import (
	"encoding/json"
	"io/fs"
)

func ReadDir(dir string) ([]fs.FileInfo, error) {
	switch dir {
	case "/test2/":
		return []fs.FileInfo{
			&MockFI{
				name:  "testproject1",
				isDir: true,
			},
			&MockFI{
				name:  "testproject2",
				isDir: true,
			},
		}, nil
	default:
		return []fs.FileInfo{}, nil
	}
}

func Mkdir(name string, p fs.FileMode) error {
	return nil
}

type project struct {
	Name            string
	Created         string
	LastModified    string
	NumberOfEntries int
}

func ReadFile(filename string) ([]byte, error) {
	switch filename {
	case "/test2/testproject1/config.json":
		data := project{
			Name:         "",
			Created:      "2022-11-13",
			LastModified: "2022-11-20",
		}
		return json.Marshal(data)
	case "/test2/testproject2/config.json":
		data := project{
			Name:         "",
			Created:      "2022-11-14",
			LastModified: "2022-11-20",
		}
		return json.Marshal(data)
	default:
		return []byte{}, nil
	}
}
