package mock_fs

import (
	"io/fs"
	"time"
)

type MockFI struct {
	name    string
	size    int64
	mode    fs.FileMode
	modtime time.Time
	isDir   bool
}

func (m *MockFI) Name() string {
	return m.name
}

func (m *MockFI) Size() int64 {
	return m.size
}

func (m *MockFI) Mode() fs.FileMode {
	return m.mode
}

func (m *MockFI) ModTime() time.Time {
	return m.modtime
}

func (m *MockFI) IsDir() bool {
	return m.isDir
}

func (m *MockFI) Sys() any {
	return nil
}
