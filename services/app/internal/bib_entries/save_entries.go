package bib_entries

import (
	"fmt"

	"github.com/TimoSto/ThesorTeX/services/app/internal/database"
)

var (
	KeyAlreadyExistsError = fmt.Errorf("key already exists")
)

type SaveEntryData struct {
	Project    string
	InitialKey string
	Key        string
	Category   string
	Fields     []string
}

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

	return err
}
