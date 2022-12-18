package bib_categories

import "github.com/TimoSto/ThesorTeX/backend/app/internal/database"

func DeleteCategory(project string, name string, store database.ThesorTeXStore) error {
	existing, err := store.GetProjectCategories(project)
	if err != nil {
		return err
	}
	for i, e := range existing {
		if e.Name == name {
			existing = append(existing[:i], existing[i+1:]...)
			break
		}
	}

	err = store.SaveProjectCategories(project, existing)
	if err != nil {
		return err
	}

	err = SaveCategoriesToSty(store, project, existing)
	if err != nil {
		return err
	}

	return nil
}
