package conf

import "github.com/TimoSto/ThesorTeX/pkg/pathbuilder"

type Config struct {
	Port        string
	ProjectsDir string
}

var config Config

func ReadConfig() {
	pathbuilder.Init()

	config = Config{
		Port:        "8448",
		ProjectsDir: pathbuilder.GetPathFromExecRoot("projects"),
	}
}

func GetConfig() Config {
	return config
}
