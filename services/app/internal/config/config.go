package config

import (
	"io/fs"
	"strconv"

	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"github.com/TimoSto/ThesorTeX/pkg/backend/pathbuilder"
	"github.com/TimoSto/ThesorTeX/services/app/internal/project_template"
	"gopkg.in/ini.v1"
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

	iniCfg, err := ini.Load("ThesorTeX.config.ini")
	if err != nil {
		log.Error("cloud not open ini: %v", err)
		iniCfg = ini.Empty()
		iniCfg.Section("").Key("port").SetValue(cfg.Port)
		iniCfg.Section("").Key("projects_dir").SetValue(cfg.ProjectsDir)
		iniCfg.Section("").Key("open_browser").SetValue(strconv.FormatBool(cfg.OpenBrowser))

		err = iniCfg.SaveTo("ThesorTeX.config.ini")
		return cfg, err
	}

	if port := iniCfg.Section("").Key("port").String(); port != "" {
		cfg.Port = port
	}
	if dir := iniCfg.Section("").Key("projects_dir").String(); dir != "" {
		cfg.ProjectsDir = dir
	}
	if openBrowser, err := iniCfg.Section("").Key("open_browser").Bool(); err == nil {
		cfg.OpenBrowser = openBrowser
	}

	return cfg, nil
}
