package updatechecker

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"github.com/TimoSto/ThesorTeX/pkg/backend/semver"
	"github.com/TimoSto/ThesorTeX/services/app/internal/config"
)

type versions struct {
	Tool []struct {
		Name string
	}
}

//TODO: unit test this
func CheckUpdateAvailable() {
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
		config.UpdateAvailable = available.ToString()
	}
}
