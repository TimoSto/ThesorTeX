package bib_entries

import (
	"github.com/TimoSto/ThesorTeX/backend/app/internal/database"
)

func GetSortedEntries(project string, store database.ThesorTeXStore) ([]BibEntry, error) {
	entries, err := ReadEntries(project, store)
	if err != nil {
		return nil, err
	}

	entries = SortEntries(entries)

	return entries, nil
}
