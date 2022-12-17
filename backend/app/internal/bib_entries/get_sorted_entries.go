package bib_entries

import (
	"github.com/TimoSto/ThesorTeX/backend/app/internal/database"
)

func GetSortedEntries(project string, store database.ThesorTeXStore) ([]database.BibEntry, error) {
	entries, err := store.GetProjectEntries(project)
	if err != nil {
		return nil, err
	}

	entries = SortEntries(entries)

	return entries, nil
}
