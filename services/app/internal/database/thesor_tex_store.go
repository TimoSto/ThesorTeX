package database

import (
	"github.com/TimoSto/ThesorTeX/services/app/internal/bib_categories"
	"github.com/TimoSto/ThesorTeX/services/app/internal/bib_entries"
)

type ThesorTeXStore interface {
	GetProjectNames() ([]string, error)
	GetProjectData(project string) ([]bib_entries.BibEntry, []bib_categories.BibCategory)
	SaveEntries(project string, entries []bib_entries.BibEntry) error
}
