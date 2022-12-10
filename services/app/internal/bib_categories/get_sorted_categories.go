package bib_categories

import (
	"github.com/TimoSto/ThesorTeX/services/app/internal/database"
)

func GetSortedCategories(project string, store database.ThesorTeXStore) ([]database.BibCategory, error) {
	categories, err := store.GetProjectCategories(project)
	if err != nil {
		return []database.BibCategory{}, err
	}

	categories = SortCategories(categories)

	return categories, nil
}
