package conf

import "github.com/TimoSto/ThesorTeX/pkg/pathbuilder"

type Config struct {
	Port        string
	ProjectsDir string
}

var config Config

func GetConfig() Config {
	pathbuilder.Init()

	config = Config{
		Port:        "8448",
		ProjectsDir: pathbuilder.GetPathFromExecRoot("projects"),
	}
	return config
}
