package bib_entries

import (
	"fmt"

	"github.com/TimoSto/ThesorTeX/backend/app/internal/database"
	"github.com/TimoSto/ThesorTeX/backend/app/internal/project"
)

var (
	KeyAlreadyExistsError = fmt.Errorf("key already exists")
)

func SaveEntriesToProject(projectName string, store database.ThesorTeXStore, entries []database.BibEntry, initialKeys []string) (database.ProjectMetaData, error) {

	existing, err := store.GetProjectEntries(projectName)
	if err != nil {
		return database.ProjectMetaData{}, err
	}

	for i, e := range entries {
		found := false
		for j, ex := range existing {
			if initialKeys[i] == ex.Key {
				existing[j] = e
				found = true
			} else if ex.Key == e.Key {
				return database.ProjectMetaData{}, KeyAlreadyExistsError
			}
		}
		if !found {
			existing = append(existing, e)
		}
	}

	existing = SortEntries(existing)

	err = store.SaveProjectEntries(projectName, existing)
	if err != nil {
		return database.ProjectMetaData{}, err
	}

	csvFile := GenerateCsvForEntries(existing)

	err = store.WriteCSV(projectName, csvFile)
	if err != nil {
		return database.ProjectMetaData{}, err
	}

	data, err := project.UpdateProjectLastEdited(projectName, store)
	if err != nil {
		return database.ProjectMetaData{}, err
	}

	return data, nil
}
