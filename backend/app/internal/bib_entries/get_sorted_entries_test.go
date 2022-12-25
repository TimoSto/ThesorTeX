package bib_entries

import (
	"encoding/json"
	"testing"

	"github.com/TimoSto/ThesorTeX/backend/app/internal/database/fake_store"
)

import (
	"reflect"
)

type sortScenario struct {
	title    string
	entries  []BibEntry
	expected []BibEntry
}

func TestGetSortedEntries(t *testing.T) {

	scenarios := []sortScenario{
		{
			title: "numbers and letters",
			entries: []BibEntry{
				{
					Key: "test",
				},
				{
					Key: "atest",
				},
				{
					Key: "btest",
				},
				{
					Key: "aatest",
				},
			},
			expected: []BibEntry{
				{
					Key: "aatest",
				},
				{
					Key: "atest",
				},
				{
					Key: "btest",
				},
				{
					Key: "test",
				},
			},
		},
	}

	for _, s := range scenarios {
		t.Run(s.title, func(t *testing.T) {
			fake := fake_store.Store{}

			data, _ := json.Marshal(s.entries)

			fake.WriteFileInProject("", entriesJsonFile, data)

			result, err := GetSortedEntries("", &fake)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(result, s.expected) {
				t.Errorf("expected %v but got %v", s.expected, result)
			}
		})
	}
}
