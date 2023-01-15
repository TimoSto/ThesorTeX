package config

import (
	"io/fs"

	"github.com/TimoSto/ThesorTeX/pkg/backend/pathbuilder"
	"github.com/TimoSto/ThesorTeX/services/app/internal/project_template"
)

var Version string

type Config struct {
	ProjectsDir     string
	ProjectTemplate fs.FS
}

func ReadConfig() (Config, error) {
	cfg := Config{
		ProjectsDir:     pathbuilder.GetPathFromExecRoot("/projects"),
		ProjectTemplate: project_template.ProjectTemplate,
	}

	return cfg, nil
}
