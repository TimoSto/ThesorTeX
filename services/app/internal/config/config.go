package config

import (
	"io/fs"

	"github.com/TimoSto/ThesorTeX/pkg/backend/pathbuilder"
	"github.com/TimoSto/ThesorTeX/services/app/internal/project_template"
)

var Version string

type Config struct {
	Port            string
	ProjectsDir     string
	OpenBrowser     bool
	ProjectTemplate fs.FS
}

func ReadConfig() (Config, error) {
	cfg := Config{
		Port:            "8448",
		ProjectsDir:     pathbuilder.GetPathFromExecRoot("/projects"),
		OpenBrowser:     false,
		ProjectTemplate: project_template.ProjectTemplate,
	}

	return cfg, nil
}
