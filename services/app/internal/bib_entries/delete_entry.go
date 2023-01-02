package bib_entries

import (
	"github.com/TimoSto/ThesorTeX/services/app/internal/database"
)

func DeleteEntry(project string, key string, store database.ThesorTeXStore) error {
	existing, err := ReadEntries(project, store)
	if err != nil {
		return err
	}
	for i, e := range existing {
		if e.Key == key {
			existing = append(existing[:i], existing[i+1:]...)
			break
		}
	}

	err = SaveEntries(existing, project, store)
	if err != nil {
		return err
	}

	return nil
}
