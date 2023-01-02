package pathbuilder

import (
	"testing"
)

func TestGetPathFromExecRoot(t *testing.T) {
	WorkingDir = "/test"
	res := GetPathFromExecRoot("dir1/file.txt")
	if res != "/test/dir1/file.txt" {
		t.Errorf("got wrong path %s", res)
	}
}

func TestGetProjectPath(t *testing.T) {
	res := GetProjectPath("c/projects/", "test1")
	if res != "c/projects/test1" {
		t.Errorf("got wrong path %s", res)
	}
}

func TestGetPathInProject(t *testing.T) {
	res := GetPathInProject("c/projects/", "test1", "file.txt")
	if res != "c/projects/test1/file.txt" {
		t.Errorf("got wrong path %s", res)
	}
}
