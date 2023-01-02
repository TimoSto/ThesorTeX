package bib_categories

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/TimoSto/ThesorTeX/services/app/internal/database/fake_store"
)

type scenario struct {
	title      string
	categories []BibCategory
	expected   []BibCategory
}

func TestGetSortedCategories(t *testing.T) {
	scenarios := []scenario{
		{
			title: "numbers_and_letters",
			categories: []BibCategory{
				{
					Name: "test1",
				},
				{
					Name: "test3",
				},
				{
					Name: "atest1",
				},
				{
					Name: "test2",
				},
			},
			expected: []BibCategory{
				{
					Name: "atest1",
				},
				{
					Name: "test1",
				},
				{
					Name: "test2",
				},
				{
					Name: "test3",
				},
			},
		},
		{
			title: "already_correct",
			categories: []BibCategory{
				{
					Name: "test1",
				},
				{
					Name: "test2",
				},
				{
					Name: "test3",
				},
			},
			expected: []BibCategory{
				{
					Name: "test1",
				},
				{
					Name: "test2",
				},
				{
					Name: "test3",
				},
			},
		},
	}

	for _, s := range scenarios {
		t.Run(s.title, func(t *testing.T) {
			store := fake_store.Store{}

			file, _ := json.Marshal(s.categories)

			store.WriteFileInProject("", jsonFilePath, file)

			result, err := GetSortedCategories("", &store)

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(result, s.expected) {
				t.Errorf("expected %v but got %v", s.expected, result)
			}
		})
	}
}
