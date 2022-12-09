package bib_entries

import (
	"sort"
	"strings"

	"github.com/TimoSto/ThesorTeX/services/app/internal/database"
)

func SortEntries(entries []database.BibEntry) []database.BibEntry {
	sort.Slice(entries, func(i, j int) bool {
		return strings.ToLower(entries[i].Key) < strings.ToLower(entries[j].Key)
	})

	return entries
}
