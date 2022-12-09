package bib_categories

import (
	"sort"
	"strings"

	"github.com/TimoSto/ThesorTeX/services/app/internal/database"
)

func SortCategories(entries []database.BibCategory) []database.BibCategory {
	sort.Slice(entries, func(i, j int) bool {
		return strings.ToLower(entries[i].Name) < strings.ToLower(entries[j].Name)
	})

	return entries
}
