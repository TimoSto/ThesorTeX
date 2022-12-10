package bib_entries

import "testing"

import (
	"reflect"

	"github.com/TimoSto/ThesorTeX/services/app/internal/database"
)

type sortScenario struct {
	title    string
	entries  []database.BibEntry
	expected []database.BibEntry
}

func TestSortEntries(t *testing.T) {

	scenarios := []sortScenario{
		{
			title: "numbers and letters",
			entries: []database.BibEntry{
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
			expected: []database.BibEntry{
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

			result := SortEntries(s.entries)

			if !reflect.DeepEqual(result, s.expected) {
				t.Errorf("expected %v but got %v", s.expected, result)
			}
		})
	}
}
