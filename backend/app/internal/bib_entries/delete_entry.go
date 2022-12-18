package bib_entries

import "github.com/TimoSto/ThesorTeX/backend/app/internal/database"

func DeleteEntry(project string, key string, store database.ThesorTeXStore) error {
	existing, err := store.GetProjectEntries(project)
	if err != nil {
		return err
	}
	for i, e := range existing {
		if e.Key == key {
			existing = append(existing[:i], existing[i+1:]...)
			break
		}
	}

	err = store.SaveProjectEntries(project, existing)
	if err != nil {
		return err
	}

	csvFile := GenerateCsvForEntries(existing)

	err = store.WriteCSV(project, csvFile)
	if err != nil {
		return err
	}

	return nil
}
