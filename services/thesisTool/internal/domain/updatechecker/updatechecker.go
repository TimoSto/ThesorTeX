package updatechecker

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"github.com/TimoSto/ThesorTeX/pkg/backend/semver"
	"github.com/TimoSto/ThesorTeX/services/thesisTool/internal/config"
)

type versions struct {
	Tool []struct {
		Name string
	}
}

//TODO: unit test this
//TODO: use go client and then add timeout
func CheckUpdateAvailable() {
	log.Info("Checking for updates...")
	resp, err := http.Get("https://thesortex.com/versions")
	if err != nil {
		log.Error("could not read versions from remote: %v", err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("could not read versions from remote: %v", err)
		return
	}

	var v versions
	err = json.Unmarshal(body, &v)
	if err != nil {
		log.Error("could not read versions from remote: %v", err)
		return
	}

	current, err := semver.ParseString(config.Version)
	if err != nil {
		log.Error("could not convert version from config package: %v", err)
		return
	}
	available, err := semver.ParseString(v.Tool[0].Name)
	if err != nil {
		log.Error("could not convert latest version from remote: %v", err)
		return
	}

	if semver.Compare(current, available) == 1 {
		log.Info("A new version is available: %s", available.ToString())
		config.Cfg.UpdateAvailable = available.ToString()
	}
}
