package bib_categories

import (
	"sort"
	"strings"
)

func SortCategories(entries []BibCategory) []BibCategory {
	sort.Slice(entries, func(i, j int) bool {
		return strings.ToLower(entries[i].Name) < strings.ToLower(entries[j].Name)
	})

	return entries
}
