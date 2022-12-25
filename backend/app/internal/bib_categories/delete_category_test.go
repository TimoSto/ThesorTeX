package bib_categories

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/TimoSto/ThesorTeX/backend/app/internal/database/fake_store"
)

func TestDeleteCategory(t *testing.T) {
	scenarios := []struct {
		title    string
		initial  []BibCategory
		deleted  string
		expected []BibCategory
	}{
		{
			title: "one category to be deleted",
			initial: []BibCategory{
				{
					Name: "test1",
				},
			},
			deleted:  "test1",
			expected: []BibCategory{},
		},
		{
			title: "three categories, one to be deleted",
			initial: []BibCategory{
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
			deleted: "test2",
			expected: []BibCategory{
				{
					Name: "test1",
				},
				{
					Name: "test3",
				},
			},
		},
		{
			title: "three categories, deleted is not present",
			initial: []BibCategory{
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
			deleted: "test4",
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
			fake := fake_store.Store{}
			file, _ := json.Marshal(s.initial)
			fake.WriteFileInProject("", jsonFilePath, file)

			err := DeleteCategory("", s.deleted, &fake)

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			resultFile, _ := fake.ReadFileInProject("", jsonFilePath)
			var result []BibCategory
			json.Unmarshal(resultFile, &result)

			if !reflect.DeepEqual(result, s.expected) {
				t.Errorf("got: %v, want: %v", result, s.expected)
			}
		})
	}
}
