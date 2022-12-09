package bib_entries

import (
	"encoding/json"
	"fmt"
	"io/fs"

	"github.com/TimoSto/ThesorTeX/pkg/pathbuilder"
	"github.com/TimoSto/ThesorTeX/services/app/conf"
)

var (
	KeyAlreadyExistsError = fmt.Errorf("key already exists")
)

func SaveEntriesToProject(project string, readFile func(string) ([]byte, error), writeFile func(string, []byte, fs.FileMode) error, config conf.Config, entries []BibEntry, initialKeys []string) error {

	existing, err := ReadEntriesOfProject(project, readFile, config)
	if err != nil {
		return err
	}

	for i, e := range entries {
		found := false
		for j, ex := range existing {
			if initialKeys[i] == ex.Key {
				existing[j] = ex
				found = true
			} else if ex.Key == e.Key {
				return KeyAlreadyExistsError
			}
		}
		if !found {
			existing = append(existing, e)
		}
	}

	data, err := json.Marshal(existing)
	if err != nil {
		return err
	}
	err = writeFile(pathbuilder.GetPathInProject(config.ProjectsDir, project, "bib_entries.json"), data, 0777)

	return err
}
