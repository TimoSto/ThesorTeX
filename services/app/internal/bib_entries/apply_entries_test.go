package bib_entries

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/TimoSto/ThesorTeX/services/app/internal/database/fake_store"
)

type scenario struct {
	title       string
	initial     []BibEntry
	saveObj     []BibEntry
	initialKeys []string //todo refactor this stuff
	expected    []BibEntry
}

func TestSaveEntriesToProject(t *testing.T) {
	scenarios := []scenario{
		{
			title:   "on empty project",
			initial: []BibEntry{},
			saveObj: []BibEntry{
				{
					Key: "test1",
				},
			},
			initialKeys: []string{""},
			expected: []BibEntry{
				{
					Key: "test1",
				},
			},
		},
		{
			title: "with one entry",
			initial: []BibEntry{
				{
					Key: "test1",
				},
			},
			initialKeys: []string{""},
			saveObj: []BibEntry{
				{
					Key:      "test2",
					Category: "test",
				},
			},
			expected: []BibEntry{
				{
					Key: "test1",
				},
				{
					Key:      "test2",
					Category: "test",
				},
			},
		},
		{
			title: "with one entry to override",
			initial: []BibEntry{
				{
					Key: "test1",
				},
			},
			initialKeys: []string{"test1"},
			saveObj: []BibEntry{
				{
					Key:      "test1",
					Category: "test2",
				},
			},
			expected: []BibEntry{
				{
					Key:      "test1",
					Category: "test2",
				},
			},
		},
		{
			title: "with one entry to override and one entry to save",
			initial: []BibEntry{
				{
					Key: "test1",
				},
			},
			initialKeys: []string{"", "test1"},
			saveObj: []BibEntry{
				{
					Key:      "test2",
					Category: "test",
				},
				{
					Key:      "test1",
					Category: "test",
				},
			},
			expected: []BibEntry{
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
		{
			title: "with one entry to override and one entry to save (different order)",
			initial: []BibEntry{
				{
					Key: "test1",
				},
			},
			initialKeys: []string{"test1", ""},
			saveObj: []BibEntry{
				{
					Key:      "test1",
					Category: "test1",
				},
				{
					Key:      "test2",
					Category: "test2",
				},
			},
			expected: []BibEntry{
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
	}

	for _, s := range scenarios {
		t.Run(s.title, func(t *testing.T) {
			store := fake_store.Store{
				Files: make(map[string][]byte),
			}

			data, _ := json.Marshal(s.initial)

			store.WriteFileInProject(s.title, entriesJsonFile, data)

			_, err := ApplyEntries(s.title, &store, s.saveObj, s.initialKeys)
			if err != nil {
				t.Errorf("unexpected error %v", err)
				return
			}

			gotFile, _ := store.ReadFileInProject(s.title, entriesJsonFile)
			var got []BibEntry
			json.Unmarshal(gotFile, &got)
			if !reflect.DeepEqual(got, s.expected) {
				t.Errorf("got %v but want %v", got, s.expected)
			}
		})
	}
}
