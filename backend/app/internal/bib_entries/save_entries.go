package bib_entries

import (
	"fmt"

	"github.com/TimoSto/ThesorTeX/backend/app/internal/database"
)

var (
	KeyAlreadyExistsError = fmt.Errorf("key already exists")
)

func SaveEntriesToProject(project string, store database.ThesorTeXStore, entries []database.BibEntry, initialKeys []string) error {

	existing, err := store.GetProjectEntries(project)
	if err != nil {
		return err
	}

	for i, e := range entries {
		found := false
		for j, ex := range existing {
			if initialKeys[i] == ex.Key {
				existing[j] = e
				found = true
			} else if ex.Key == e.Key {
				return KeyAlreadyExistsError
			}
		}
		if !found {
			existing = append(existing, e)
		}
	}

	existing = SortEntries(existing)

	err = store.SaveProjectEntries(project, existing)
	if err != nil {
		return err
	}

	csvFile := GenerateCsvForEntries(existing)

	err = store.WriteCSV(project, csvFile)
	if err != nil {
		return err
	}

	return nil
}