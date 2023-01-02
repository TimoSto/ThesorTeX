package local_store

import (
	"github.com/TimoSto/ThesorTeX/services/app/internal/conf"
)

type Store struct {
	Config conf.Config
}

const (
	projectEntriesFile = "data/bib_entries.json"

	projectCategoriesFile = "data/bib_categories.json"

	entriesCsvFile = "bib_entries.csv"

	bibliographyStyFile = "styPackages/bibliography.sty"
)
