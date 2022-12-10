package bib_categories

import (
	"reflect"
	"testing"

	"github.com/TimoSto/ThesorTeX/services/app/internal/database"
)

type scenario struct {
	title      string
	categories []database.BibCategory
	expected   []database.BibCategory
}

func TestGetSortedCategories(t *testing.T) {
	scenarios := []scenario{
		{
			title: "numbers_and_letters",
			categories: []database.BibCategory{
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
			expected: []database.BibCategory{
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
			categories: []database.BibCategory{
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
			expected: []database.BibCategory{
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

			result := SortCategories(s.categories)

			if !reflect.DeepEqual(result, s.expected) {
				t.Errorf("expected %v but got %v", s.expected, result)
			}
		})
	}
}
