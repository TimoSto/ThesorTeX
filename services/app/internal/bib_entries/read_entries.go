package bib_entries

import (
	"encoding/json"
	"fmt"

	"github.com/TimoSto/ThesorTeX/pkg/pathbuilder"
	"github.com/TimoSto/ThesorTeX/services/app/conf"
)

func ReadEntriesOfProject(project string, readFile func(string) ([]byte, error), config conf.Config) ([]BibEntry, error) {
	if len(project) == 0 {
		return nil, fmt.Errorf("no project name provided")
	}
	file, err := readFile(pathbuilder.GetPathInProject(config.ProjectsDir, project, "bib_entries.json"))
	if err != nil {
		return nil, err
	}

	var bibEntries []BibEntry
	err = json.Unmarshal(file, &bibEntries)
	if err != nil {
		return nil, err
	}

	return bibEntries, nil
}
