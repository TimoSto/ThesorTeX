package database

import (
	"github.com/TimoSto/ThesorTeX/services/app/internal/bib_categories"
)

type ProjectMetaData struct {
	Name            string
	Created         string
	LastModified    string
	NumberOfEntries int
}

type ProjectData struct {
	Entries    []BibEntry
	Categories []bib_categories.BibCategory
}
