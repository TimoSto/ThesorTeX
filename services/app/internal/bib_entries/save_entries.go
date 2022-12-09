package bib_entries

import (
	"fmt"

	"github.com/TimoSto/ThesorTeX/services/app/internal/database"
)

var (
	KeyAlreadyExistsError = fmt.Errorf("key already exists")
)

type SaveEntryData struct {
	Project    string
	InitialKey string
	Key        string
	Category   string
	Fields     []string
}

func SaveEntriesToProject(project string, store database.ThesorTeXStore, entries []database.BibEntry, initialKeys []string) error {

	projectData, err := store.GetProjectData(project)
	if err != nil {
		return err
	}

	for i, e := range entries {
		found := false
		for j, ex := range projectData.Entries {
			if initialKeys[i] == ex.Key {
				projectData.Entries[j] = e
				found = true
			} else if ex.Key == e.Key {
				return KeyAlreadyExistsError
			}
		}
		if !found {
			projectData.Entries = append(projectData.Entries, e)
		}
	}

	fmt.Println(projectData.Entries[0].Fields[0])

	projectData.Entries = SortEntries(projectData.Entries)

	err = store.SaveProjectData(project, projectData)

	return err
}
