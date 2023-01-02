package bib_categories

import (
	"sort"
	"strings"

	"github.com/TimoSto/ThesorTeX/services/app/internal/database"
)

func GetSortedCategories(project string, store database.ThesorTeXStore) ([]BibCategory, error) {
	categories, err := ReadCategories(project, store)
	if err != nil {
		return []BibCategory{}, err
	}

	sort.Slice(categories, func(i, j int) bool {
		return strings.ToLower(categories[i].Name) < strings.ToLower(categories[j].Name)
	})

	return categories, nil
}
