package bib_entries

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/TimoSto/ThesorTeX/services/app/internal/database/fake_store"
)

func TestDeleteEntry(t *testing.T) {
	scenarios := []struct {
		title    string
		initial  []BibEntry
		deleted  string
		expected []BibEntry
	}{
		{
			title: "one Entry to be deleted",
			initial: []BibEntry{
				{
					Key: "test1",
				},
			},
			deleted:  "test1",
			expected: []BibEntry{},
		},
		{
			title: "three categories, one to be deleted",
			initial: []BibEntry{
				{
					Key: "test1",
				},
				{
					Key: "test2",
				},
				{
					Key: "test3",
				},
			},
			deleted: "test2",
			expected: []BibEntry{
				{
					Key: "test1",
				},
				{
					Key: "test3",
				},
			},
		},
		{
			title: "three categories, deleted is not present",
			initial: []BibEntry{
				{
					Key: "test1",
				},
				{
					Key: "test2",
				},
				{
					Key: "test3",
				},
			},
			deleted: "test4",
			expected: []BibEntry{
				{
					Key: "test1",
				},
				{
					Key: "test2",
				},
				{
					Key: "test3",
				},
			},
		},
	}

	for _, s := range scenarios {
		t.Run(s.title, func(t *testing.T) {
			fake := fake_store.Store{}
			file, _ := json.Marshal(s.initial)
			fake.WriteFileInProject("", entriesJsonFile, file)

			err := DeleteEntry("", s.deleted, &fake)

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			resultFile, _ := fake.ReadFileInProject("", entriesJsonFile)
			var result []BibEntry
			json.Unmarshal(resultFile, &result)

			if !reflect.DeepEqual(result, s.expected) {
				t.Errorf("got: %v, want: %v", result, s.expected)
			}
		})
	}
}
