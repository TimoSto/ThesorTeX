package bib_categories

import "github.com/TimoSto/ThesorTeX/services/app/internal/database"

func SaveCategory(store database.ThesorTeXStore, project string, name string, initialName string, citaviCategory string, citaviFilter []string, bibFields []database.Field, citeFields []database.Field) error {
	existing, err := store.GetProjectCategories(project)
	if err != nil {
		return err
	}

	found := false

	for i, e := range existing {
		if e.Name == initialName {
			existing[i].Name = name
			existing[i].CitaviCategory = citaviCategory
			existing[i].CitaviFilters = citaviFilter
			existing[i].Fields = bibFields
			existing[i].CiteFields = citeFields

			found = true
			break
		}
	}

	if !found {
		existing = append(existing, database.BibCategory{
			Name:           name,
			CitaviCategory: citaviCategory,
			CitaviFilters:  citaviFilter,
			Fields:         bibFields,
			CiteFields:     citeFields,
		})
	}

	err = store.SaveProjectCategories(project, existing)
	if err != nil {
		return err
	}

	return SaveCategoriesToSty(store, project, existing)
}
