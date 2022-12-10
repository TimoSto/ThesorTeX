package bib_entries

import (
	"github.com/TimoSto/ThesorTeX/services/app/internal/database"
)

func GetSortedEntries(project string, store database.ThesorTeXStore) ([]database.BibEntry, error) {
	entries, err := store.GetProjectEntries(project)
	if err != nil {
		return nil, err
	}

	return entries, nil
}
