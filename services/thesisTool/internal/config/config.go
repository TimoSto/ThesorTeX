package config

import (
	"io/fs"
	"strconv"

	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"github.com/TimoSto/ThesorTeX/pkg/backend/pathbuilder"
	"github.com/TimoSto/ThesorTeX/pkg/backend/project_template"
	"gopkg.in/ini.v1"
)

var Version string

type Config struct {
	Port            string
	ProjectsDir     string
	OpenBrowser     bool
	ProjectTemplate fs.FS
	UpdateAvailable string
	LogLevel        string
}

var Cfg Config

func ReadConfig() (Config, error) {
	cfg := Config{
		Port:            "8448",
		ProjectsDir:     pathbuilder.GetPathFromExecRoot("/projects"),
		OpenBrowser:     false,
		ProjectTemplate: project_template.ProjectTemplate,
		LogLevel:        "INFO",
	}

	iniCfg, err := ini.Load(pathbuilder.GetPathFromExecRoot("ThesorTeX.config.ini"))
	if err != nil {
		log.Info("could not open ini: %v", err)

		err = SaveConfig(cfg)

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
	if logLevel := iniCfg.Section("").Key("logLevel").String(); logLevel != "" {
		cfg.LogLevel = logLevel
	}

	Cfg = cfg

	return cfg, nil
}

func SaveConfig(cfg Config) error {
	iniCfg := ini.Empty()
	iniCfg.Section("").Key("port").SetValue(cfg.Port)
	iniCfg.Section("").Key("projects_dir").SetValue(cfg.ProjectsDir)
	iniCfg.Section("").Key("open_browser").SetValue(strconv.FormatBool(cfg.OpenBrowser))
	iniCfg.Section("").Key("logLevel").SetValue(cfg.LogLevel)

	err := iniCfg.SaveTo(pathbuilder.GetPathFromExecRoot("ThesorTeX.config.ini"))

	if err != nil {
		return err
	}

	Cfg = cfg

	return nil
}
