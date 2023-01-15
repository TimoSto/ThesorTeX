package config

import "github.com/TimoSto/ThesorTeX/pkg/backend/pathbuilder"

var Version string

type Config struct {
	ProjectsDir string
}

func ReadConfig() (Config, error) {
	cfg := Config{
		ProjectsDir: pathbuilder.GetPathFromExecRoot("/projects"),
	}

	return cfg, nil
}
