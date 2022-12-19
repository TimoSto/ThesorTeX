package bib_categories

import (
	"github.com/TimoSto/ThesorTeX/backend/app/internal/database"
	"github.com/TimoSto/ThesorTeX/backend/app/internal/project"
)

func SaveCategory(store database.ThesorTeXStore, projectName string, name string, initialName string, citaviCategory string, citaviFilter []string, bibFields []database.Field, citeFields []database.Field) (database.ProjectMetaData, error) {
	existing, err := store.GetProjectCategories(projectName)
	if err != nil {
		return database.ProjectMetaData{}, err
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

	err = store.SaveProjectCategories(projectName, existing)
	if err != nil {
		return database.ProjectMetaData{}, err
	}

	err = SaveCategoriesToSty(store, projectName, existing)
	if err != nil {
		return database.ProjectMetaData{}, err
	}

	data, err := project.UpdateProjectLastEdited(projectName, store)
	if err != nil {
		return database.ProjectMetaData{}, err
	}

	return data, nil
}
