package bib_entries

import (
	"encoding/json"

	"github.com/TimoSto/ThesorTeX/backend/app/internal/database"
)

func ReadEntries(project string, store database.ThesorTeXStore) ([]BibEntry, error) {
	file, err := store.ReadFileInProject(project, entriesJsonFile)
	if err != nil {
		return nil, err
	}

	var entries []BibEntry
	err = json.Unmarshal(file, &entries)
	if err != nil {
		return nil, err
	}

	return entries, nil
}
