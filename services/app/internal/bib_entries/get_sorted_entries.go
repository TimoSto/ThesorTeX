package bib_entries

import (
	"sort"
	"strings"

	"github.com/TimoSto/ThesorTeX/services/app/internal/database"
)

func GetSortedEntries(project string, store database.ThesorTeXStore) ([]BibEntry, error) {
	entries, err := ReadEntries(project, store)
	if err != nil {
		return nil, err
	}

	sort.Slice(entries, func(i, j int) bool {
		return strings.ToLower(entries[i].Key) < strings.ToLower(entries[j].Key)
	})

	return entries, nil
}
