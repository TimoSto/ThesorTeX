package bib_categories

import (
	"fmt"

	"github.com/TimoSto/ThesorTeX/services/app/internal/database"
	project2 "github.com/TimoSto/ThesorTeX/services/app/internal/project"
)

func SaveCategory(store database.ThesorTeXStore, projectName string, name string, initialName string, citaviCategory string, citaviFilter []string, bibFields []Field, citeFields []Field) (project2.ProjectMetaData, error) {
	fmt.Println("gothere1")
	existing, err := ReadCategories(projectName, store)
	if err != nil {
		return project2.ProjectMetaData{}, err
	}
	fmt.Println("gothere2")

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
		existing = append(existing, BibCategory{
			Name:           name,
			CitaviCategory: citaviCategory,
			CitaviFilters:  citaviFilter,
			Fields:         bibFields,
			CiteFields:     citeFields,
		})
	}

	err = SaveCategoriesToJSON(projectName, existing, store)
	if err != nil {
		return project2.ProjectMetaData{}, err
	}

	err = SaveCategoriesToSty(store, projectName, existing)
	if err != nil {
		return project2.ProjectMetaData{}, err
	}

	data, err := project2.UpdateProjectLastEdited(projectName, store)
	if err != nil {
		return project2.ProjectMetaData{}, err
	}

	return data, nil
}
