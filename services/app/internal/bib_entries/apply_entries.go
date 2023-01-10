package bib_entries

import (
	"fmt"

	"github.com/TimoSto/ThesorTeX/services/app/internal/database"
	project2 "github.com/TimoSto/ThesorTeX/services/app/internal/project"
)

var (
	KeyAlreadyExistsError = fmt.Errorf("key already exists")
)

func ApplyEntries(projectName string, store database.ThesorTeXStore, entries []BibEntry, initialKeys []string) (project2.ProjectMetaData, error) {

	existing, err := ReadEntries(projectName, store)
	if err != nil {
		return project2.ProjectMetaData{}, err
	}

	if initialKeys == nil {
		for _, e := range entries {
			initialKeys = append(initialKeys, e.Key)
		}
	}

	for i, e := range entries {
		found := false
		for j, ex := range existing {
			if initialKeys[i] == ex.Key {
				existing[j] = e
				found = true
			} else if ex.Key == e.Key {
				return project2.ProjectMetaData{}, KeyAlreadyExistsError
			}
		}
		if !found {
			existing = append(existing, e)
		}
	}

	//TODO: is it ok to leave out sorting here or only do it here

	err = SaveEntries(existing, projectName, store)
	if err != nil {
		return project2.ProjectMetaData{}, err
	}

	data, err := project2.UpdateProjectLastEdited(projectName, store)
	if err != nil {
		return project2.ProjectMetaData{}, err
	}

	return data, nil
}
