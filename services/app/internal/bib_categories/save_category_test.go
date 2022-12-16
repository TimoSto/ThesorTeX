package bib_categories

import (
	"reflect"
	"testing"

	"github.com/TimoSto/ThesorTeX/services/app/internal/database"
	"github.com/TimoSto/ThesorTeX/services/app/internal/database/fake_store"
)

func TestSaveCategory(t *testing.T) {
	tcs := []struct {
		title          string
		initial        [][]database.BibCategory
		name           string
		initialName    string
		citaviCategory string
		citaviFilter   []string
		bibFields      []database.Field
		citeFields     []database.Field
		expected       []database.BibCategory
	}{
		{
			title: "save on empty",
			initial: [][]database.BibCategory{
				{},
			},
			name: "test",
			bibFields: []database.Field{
				{
					Field:            "test1",
					Italic:           true,
					Prefix:           "(",
					Suffix:           ") ",
					TexValue:         false,
					CitaviAttributes: nil,
				},
			},
			expected: []database.BibCategory{
				{
					Name:           "test",
					CitaviCategory: "",
					CitaviFilters:  nil,
					Fields: []database.Field{
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
			initial: [][]database.BibCategory{
				{
					{
						Name:           "test",
						CitaviCategory: "",
						CitaviFilters:  nil,
						Fields: []database.Field{
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
			name:        "test",
			initialName: "test",
			bibFields: []database.Field{
				{
					Field:            "test1",
					Italic:           false,
					Prefix:           "(",
					Suffix:           ") ",
					TexValue:         false,
					CitaviAttributes: nil,
				},
			},
			expected: []database.BibCategory{
				{
					Name:           "test",
					CitaviCategory: "",
					CitaviFilters:  nil,
					Fields: []database.Field{
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
			initial: [][]database.BibCategory{
				{
					{
						Name:           "test",
						CitaviCategory: "",
						CitaviFilters:  nil,
						Fields: []database.Field{
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
			name:        "test2",
			initialName: "test",
			bibFields: []database.Field{
				{
					Field:            "test1",
					Italic:           false,
					Prefix:           "(",
					Suffix:           ") ",
					TexValue:         false,
					CitaviAttributes: nil,
				},
			},
			expected: []database.BibCategory{
				{
					Name:           "test2",
					CitaviCategory: "",
					CitaviFilters:  nil,
					Fields: []database.Field{
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
			store := fake_store.Store{
				Entries:    nil,
				Categories: tc.initial,
				ProjectMeta: []database.ProjectMetaData{
					{
						Name: "testproject",
					},
				},
			}

			err := SaveCategory(&store, "testproject", tc.name, tc.initialName, tc.citaviCategory, tc.citaviFilter, tc.bibFields, tc.citeFields)

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			got, _ := store.GetProjectCategories("testproject")
			if !reflect.DeepEqual(tc.expected, got) {
				t.Errorf("got categories: %v, want: %v", got, tc.expected)
			}

		})
	}
}
