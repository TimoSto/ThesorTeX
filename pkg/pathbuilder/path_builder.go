package pathbuilder

import (
	"fmt"
	"os"
	"path/filepath"
)

var WorkingDir string

func Init() {
	executablePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error reading current dir: " + err.Error())
		fmt.Println("Maybe there are some wrong permissions")
		os.Exit(1)
	}

	WorkingDir = filepath.Dir(executablePath)
}

func GetPathInProject(projectsDir string, project string, file string) string {
	return filepath.Join(projectsDir, project, file)
}

func GetProjectPath(projectsDir string, project string) string {
	return filepath.Join(projectsDir, project)
}

func GetPathFromExecRoot(path string) string {
	return filepath.Join(WorkingDir, path)
}
