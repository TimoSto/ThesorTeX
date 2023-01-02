package bib_categories

import (
	"encoding/json"

	"github.com/TimoSto/ThesorTeX/services/app/internal/database"
)

func SaveCategoriesToJSON(project string, categories []BibCategory, store database.ThesorTeXStore) error {
	data, err := json.MarshalIndent(categories, "", "\t")
	if err != nil {
		return err
	}

	err = store.WriteFileInProject(project, jsonFilePath, data)
	if err != nil {
		return err
	}

	return nil
}
