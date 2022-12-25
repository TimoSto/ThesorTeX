package bib_entries

import (
	"fmt"

	"github.com/TimoSto/ThesorTeX/backend/app/internal/database"
	"github.com/TimoSto/ThesorTeX/backend/app/internal/project"
)

var (
	KeyAlreadyExistsError = fmt.Errorf("key already exists")
)

func ApplyEntries(projectName string, store database.ThesorTeXStore, entries []BibEntry, initialKeys []string) (project.ProjectMetaData, error) {

	existing, err := ReadEntries(projectName, store)
	if err != nil {
		return project.ProjectMetaData{}, err
	}

	for i, e := range entries {
		found := false
		for j, ex := range existing {
			if initialKeys[i] == ex.Key {
				existing[j] = e
				found = true
			} else if ex.Key == e.Key {
				return project.ProjectMetaData{}, KeyAlreadyExistsError
			}
		}
		if !found {
			existing = append(existing, e)
		}
	}

	//TODO: is it ok to leave out sorting here or only do it here

	err = SaveEntries(existing, projectName, store)
	if err != nil {
		return project.ProjectMetaData{}, err
	}

	data, err := project.UpdateProjectLastEdited(projectName, store)
	if err != nil {
		return project.ProjectMetaData{}, err
	}

	return data, nil
}
