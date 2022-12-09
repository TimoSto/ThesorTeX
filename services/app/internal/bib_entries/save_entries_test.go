package bib_entries

import (
	"testing"

	"github.com/TimoSto/ThesorTeX/services/app/internal/database"
)

type scenario struct {
	title    string
	initial  []database.BibEntry
	saveObj  SaveEntryData
	expected []database.BibEntry
}

func TestSaveEntriesToProject(t *testing.T) {
	scenarios := []scenario{
		{
			title:    "on empty project",
			initial:  nil,
			saveObj:  SaveEntryData{},
			expected: nil,
		},
		{
			title:    "with one entry",
			initial:  nil,
			saveObj:  SaveEntryData{},
			expected: nil,
		},
		{
			title:    "with one entry to override",
			initial:  nil,
			saveObj:  SaveEntryData{},
			expected: nil,
		},
		{
			title:    "with one entry and two entries to add",
			initial:  nil,
			saveObj:  SaveEntryData{},
			expected: nil,
		},
		{
			title:    "with one entry to override and one entry to save",
			initial:  nil,
			saveObj:  SaveEntryData{},
			expected: nil,
		},
	}

	for _, s := range scenarios {
		t.Run(s.title, func(t *testing.T) {
			//TODO: add when db is created
		})
	}
}
