package bib_categories

import (
	"encoding/json"

	"github.com/TimoSto/ThesorTeX/pkg/pathbuilder"
	"github.com/TimoSto/ThesorTeX/services/app/conf"
	"github.com/TimoSto/ThesorTeX/services/app/internal/database"
)

func ReadCategoriesOfProject(project string, readFile func(string) ([]byte, error), config conf.Config) ([]database.BibCategory, error) {
	file, err := readFile(pathbuilder.GetPathInProject(config.ProjectsDir, project, "bib_categories.json"))
	if err != nil {
		return []database.BibCategory{}, err
	}
	var categories []database.BibCategory

	err = json.Unmarshal(file, &categories)
	if err != nil {
		return []database.BibCategory{}, err
	}

	return categories, nil
}
