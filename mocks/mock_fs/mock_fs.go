package mock_fs

import "io/fs"

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
