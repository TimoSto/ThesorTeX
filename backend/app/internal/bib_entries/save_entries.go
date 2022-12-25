package bib_entries

import (
	"encoding/json"

	"github.com/TimoSto/ThesorTeX/backend/app/internal/database"
)

func SaveEntries(entries []BibEntry, project string, store database.ThesorTeXStore) error {
	file, err := json.Marshal(entries)
	if err != nil {
		return err
	}

	err = store.WriteFileInProject(project, entriesJsonFile, file)
	if err != nil {
		return err
	}

	csvFile := GenerateCsvForEntries(entries)

	err = store.WriteFileInProject(project, entriesCsvFile, []byte(csvFile))
	if err != nil {
		return err
	}

	return nil
}
