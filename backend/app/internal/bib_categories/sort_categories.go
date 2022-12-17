package bib_categories

import (
	"sort"
	"strings"

	"github.com/TimoSto/ThesorTeX/backend/app/internal/database"
)

func SortCategories(categories []database.BibCategory) []database.BibCategory {
	sort.Slice(categories, func(i, j int) bool {
		return strings.ToLower(categories[i].Name) < strings.ToLower(categories[j].Name)
	})

	return categories
}
