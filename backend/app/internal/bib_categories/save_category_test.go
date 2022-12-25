package bib_categories

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/TimoSto/ThesorTeX/backend/app/internal/database/fake_store"
)

func TestSaveCategory(t *testing.T) {
	tcs := []struct {
		title          string
		initial        []BibCategory
		name           string
		initialName    string
		citaviCategory string
		citaviFilter   []string
		bibFields      []Field
		citeFields     []Field
		expected       []BibCategory
	}{
		{
			title: "save on empty",
			initial: []BibCategory{
				{},
			},
			name: "test",
			bibFields: []Field{
				{
					Field:            "test1",
					Italic:           true,
					Prefix:           "(",
					Suffix:           ") ",
					TexValue:         false,
					CitaviAttributes: nil,
				},
			},
			expected: []BibCategory{
				{
					Name:           "test",
					CitaviCategory: "",
					CitaviFilters:  nil,
					Fields: []Field{
						{
							Field:            "test1",
							Italic:           true,
							Prefix:           "(",
							Suffix:           ") ",
							TexValue:         false,
							CitaviAttributes: nil,
						},
					},
					CiteFields: nil,
				},
			},
		},
		{
			title: "override existing",
			initial: []BibCategory{
				{
					Name:           "test",
					CitaviCategory: "",
					CitaviFilters:  nil,
					Fields: []Field{
						{
							Field:            "test1",
							Italic:           true,
							Prefix:           "(",
							Suffix:           ") ",
							TexValue:         false,
							CitaviAttributes: nil,
						},
					},
					CiteFields: nil,
				},
			},
			name:        "test",
			initialName: "test",
			bibFields: []Field{
				{
					Field:            "test1",
					Italic:           false,
					Prefix:           "(",
					Suffix:           ") ",
					TexValue:         false,
					CitaviAttributes: nil,
				},
			},
			expected: []BibCategory{
				{
					Name:           "test",
					CitaviCategory: "",
					CitaviFilters:  nil,
					Fields: []Field{
						{
							Field:            "test1",
							Italic:           false,
							Prefix:           "(",
							Suffix:           ") ",
							TexValue:         false,
							CitaviAttributes: nil,
						},
					},
					CiteFields: nil,
				},
			},
		},
		{
			title: "override and rename existing",
			initial: []BibCategory{
				{
					Name:           "test",
					CitaviCategory: "",
					CitaviFilters:  nil,
					Fields: []Field{
						{
							Field:            "test1",
							Italic:           true,
							Prefix:           "(",
							Suffix:           ") ",
							TexValue:         false,
							CitaviAttributes: nil,
						},
					},
					CiteFields: nil,
				},
			},
			name:        "test2",
			initialName: "test",
			bibFields: []Field{
				{
					Field:            "test1",
					Italic:           false,
					Prefix:           "(",
					Suffix:           ") ",
					TexValue:         false,
					CitaviAttributes: nil,
				},
			},
			expected: []BibCategory{
				{
					Name:           "test2",
					CitaviCategory: "",
					CitaviFilters:  nil,
					Fields: []Field{
						{
							Field:            "test1",
							Italic:           false,
							Prefix:           "(",
							Suffix:           ") ",
							TexValue:         false,
							CitaviAttributes: nil,
						},
					},
					CiteFields: nil,
				},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.title, func(t *testing.T) {
			store := fake_store.Store{}

			file, _ := json.Marshal(tc.initial)
			store.WriteFileInProject("testproject", jsonFilePath, file)

			_, err := SaveCategory(&store, "testproject", tc.name, tc.initialName, tc.citaviCategory, tc.citaviFilter, tc.bibFields, tc.citeFields)

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			gotFile, _ := store.ReadFileInProject("testproject", jsonFilePath)
			var got []BibCategory
			json.Unmarshal(gotFile, &got)
			if !reflect.DeepEqual(tc.expected, got) {
				t.Errorf("got categories: %v, want: %v", got, tc.expected)
			}

		})
	}
}
