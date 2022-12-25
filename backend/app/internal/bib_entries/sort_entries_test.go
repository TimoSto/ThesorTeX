package bib_entries

import "testing"

import (
	"reflect"
)

type sortScenario struct {
	title    string
	entries  []BibEntry
	expected []BibEntry
}

func TestSortEntries(t *testing.T) {

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

			result := SortEntries(s.entries)

			if !reflect.DeepEqual(result, s.expected) {
				t.Errorf("expected %v but got %v", s.expected, result)
			}
		})
	}
}
