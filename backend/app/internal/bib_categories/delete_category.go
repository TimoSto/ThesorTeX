package bib_categories

import (
	"github.com/TimoSto/ThesorTeX/backend/app/internal/database"
)

func DeleteCategory(project string, name string, store database.ThesorTeXStore) error {
	categories, err := ReadCategories(project, store)
	if err != nil {
		return err
	}

	for i, e := range categories {
		if e.Name == name {
			categories = append(categories[:i], categories[i+1:]...)
			break
		}
	}

	err = SaveCategoriesToJSON(project, categories, store)
	if err != nil {
		return err
	}

	err = SaveCategoriesToSty(store, project, categories)
	if err != nil {
		return err
	}

	return nil
}
