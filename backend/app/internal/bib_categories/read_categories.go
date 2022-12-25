package bib_categories

import (
	"encoding/json"

	"github.com/TimoSto/ThesorTeX/backend/app/internal/database"
)

func ReadCategories(project string, store database.ThesorTeXStore) ([]BibCategory, error) {
	jsonFile, err := store.ReadFileInProject(project, jsonFilePath)
	if err != nil {
		return nil, err
	}

	var categories []BibCategory
	err = json.Unmarshal(jsonFile, &categories)
	if err != nil {
		return nil, err
	}

	return categories, nil
}
