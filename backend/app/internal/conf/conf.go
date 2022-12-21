package conf

import (
	"os"

	"github.com/TimoSto/ThesorTeX/backend/pkg/log"
	"github.com/TimoSto/ThesorTeX/backend/pkg/pathbuilder"
	"gopkg.in/ini.v1"
)

type Config struct {
	Port        string
	ProjectsDir string
}

var Version string

const (
	PORT_KEY         = "PORT"
	PROJECTS_DIR_KEY = "PROJECTS_DIR"
)

var config Config

func ReadConfig() {
	pathbuilder.Init()

	config = Config{
		Port:        "8448",
		ProjectsDir: pathbuilder.GetPathFromExecRoot("projects"),
	}

	cfg, err := ini.Load("Config.ini")

	if err != nil {
		log.Error("Reading config: %v", err)
		return
	}

	cfgPort := cfg.Section("").Key(PORT_KEY).String()
	if cfgPort != "" {
		config.Port = cfgPort
	}

	cfgProjectsDir := cfg.Section("").Key(PROJECTS_DIR_KEY).String()
	if cfgProjectsDir != "" {
		config.ProjectsDir = cfgProjectsDir
	}
}

func GetConfig() Config {
	return config
}

func WriteConfig(c Config) error {
	config = c

	cfg, err := ini.Load("Config.ini")
	if err != nil {
		_, err = os.Create("Config.ini")
		if err != nil {
			return err
		}
		cfg, err = ini.Load("Config.ini")
		if err != nil {
			return err
		}
	}

	cfg.Section("").Key(PORT_KEY).SetValue(config.Port)

	cfg.Section("").Key(PROJECTS_DIR_KEY).SetValue(config.ProjectsDir)

	return cfg.SaveTo("Config.ini")
}
