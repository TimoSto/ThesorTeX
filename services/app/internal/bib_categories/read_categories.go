package bib_categories

import (
	"encoding/json"

	"github.com/TimoSto/ThesorTeX/pkg/pathbuilder"
	"github.com/TimoSto/ThesorTeX/services/app/conf"
)

func ReadCategoriesOfProject(project string, readFile func(string) ([]byte, error), config conf.Config) ([]BibCategory, error) {
	file, err := readFile(pathbuilder.GetPathInProject(config.ProjectsDir, project, "bib_categories.json"))
	if err != nil {
		return []BibCategory{}, err
	}
	var categories []BibCategory

	err = json.Unmarshal(file, &categories)
	if err != nil {
		return []BibCategory{}, err
	}

	return categories, nil
}
