package bib_entries

import (
	"reflect"
	"testing"

	"github.com/TimoSto/ThesorTeX/backend/app/internal/database"
	"github.com/TimoSto/ThesorTeX/backend/app/internal/database/fake_store"
)

type scenario struct {
	title       string
	initial     [][]database.BibEntry
	saveObj     []database.BibEntry
	initialKeys []string //todo refactor this stuff
	expected    [][]database.BibEntry
}

func TestSaveEntriesToProject(t *testing.T) {
	scenarios := []scenario{
		{
			title: "on empty project",
			initial: [][]database.BibEntry{
				{},
			},
			saveObj: []database.BibEntry{
				{
					Key: "test1",
				},
			},
			initialKeys: []string{""},
			expected: [][]database.BibEntry{
				{
					{
						Key: "test1",
					},
				},
			},
		},
		{
			title: "with one entry",
			initial: [][]database.BibEntry{
				{
					{
						Key: "test1",
					},
				},
			},
			initialKeys: []string{""},
			saveObj: []database.BibEntry{
				{
					Key:      "test2",
					Category: "test",
				},
			},
			expected: [][]database.BibEntry{
				{
					{
						Key: "test1",
					},
					{
						Key:      "test2",
						Category: "test",
					},
				},
			},
		},
		{
			title: "with one entry to override",
			initial: [][]database.BibEntry{
				{
					{
						Key: "test1",
					},
				},
			},
			initialKeys: []string{"test1"},
			saveObj: []database.BibEntry{
				{
					Key:      "test1",
					Category: "test2",
				},
			},
			expected: [][]database.BibEntry{
				{
					{
						Key:      "test1",
						Category: "test2",
					},
				},
			},
		},
		{
			title: "with one entry to override and one entry to save",
			initial: [][]database.BibEntry{
				{
					{
						Key: "test1",
					},
				},
			},
			initialKeys: []string{"", "test1"},
			saveObj: []database.BibEntry{
				{
					Key:      "test2",
					Category: "test",
				},
				{
					Key:      "test1",
					Category: "test",
				},
			},
			expected: [][]database.BibEntry{
				{
					{
						Key:      "test1",
						Category: "test",
					},
					{
						Key:      "test2",
						Category: "test",
					},
				},
			},
		},
		{
			title: "with one entry to override and one entry to save (different order)",
			initial: [][]database.BibEntry{
				{
					{
						Key: "test1",
					},
				},
			},
			initialKeys: []string{"test1", ""},
			saveObj: []database.BibEntry{
				{
					Key:      "test1",
					Category: "test1",
				},
				{
					Key:      "test2",
					Category: "test2",
				},
			},
			expected: [][]database.BibEntry{
				{
					{
						Key:      "test1",
						Category: "test1",
					},
					{
						Key:      "test2",
						Category: "test2",
					},
				},
			},
		},
	}

	for _, s := range scenarios {
		t.Run(s.title, func(t *testing.T) {
			store := fake_store.Store{
				Entries:    s.initial,
				Categories: nil,
				ProjectMeta: []database.ProjectMetaData{
					{
						Name: s.title,
					},
				},
			}

			SaveEntriesToProject(s.title, &store, s.saveObj, s.initialKeys)

			if !reflect.DeepEqual(store.Entries, s.expected) {
				t.Errorf("expected %v but got %v", s.expected, store.Entries)
			}
		})
	}
}
