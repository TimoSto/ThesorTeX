package bib_categories

import (
	"reflect"
	"testing"

	"github.com/TimoSto/ThesorTeX/mocks/mock_fs"
	"github.com/TimoSto/ThesorTeX/services/app/conf"
	"github.com/TimoSto/ThesorTeX/services/app/internal/database"
)

func TestReadCategoriesOfProject(t *testing.T) {
	categories, err := ReadCategoriesOfProject("test_threeCategories", mock_fs.ReadFile, conf.Config{ProjectsDir: "/test"})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(categories) != 3 {
		t.Errorf("expected 3 categories, but got %v", len(categories))
	}

	expected := []database.BibCategory{
		{
			Name:                  "test1",
			CitaviType:            "",
			CitaviNecessaryFields: nil,
			Fields: []database.Field{
				{
					Field:            "t11",
					Style:            "normal",
					Prefix:           "",
					Suffix:           ", ",
					TexValue:         false,
					CitaviAttributes: nil,
				},
				{
					Field:            "t12",
					Style:            "italic",
					Prefix:           " ",
					Suffix:           ".",
					TexValue:         false,
					CitaviAttributes: nil,
				},
			},
			CiteFields: []database.Field{
				{
					Field:            "t11",
					Style:            "normal",
					Prefix:           "",
					Suffix:           ", ",
					TexValue:         false,
					CitaviAttributes: nil,
				},
			},
		},
		{
			Name:                  "test2",
			CitaviType:            "citaviTestDoi",
			CitaviNecessaryFields: []string{"doi"},
			Fields: []database.Field{
				{
					Field:            "t21",
					Style:            "normal",
					Prefix:           "",
					Suffix:           ", ",
					TexValue:         false,
					CitaviAttributes: []string{"author"},
				},
				{
					Field:            "t22",
					Style:            "italic",
					Prefix:           " ",
					Suffix:           ".",
					TexValue:         false,
					CitaviAttributes: []string{"doi"},
				},
			},
			CiteFields: []database.Field{
				{
					Field:            "t23",
					Style:            "normal",
					Prefix:           "",
					Suffix:           ", ",
					TexValue:         false,
					CitaviAttributes: nil,
				},
			},
		},
		{
			Name:                  "test3",
			CitaviType:            "citaviTest",
			CitaviNecessaryFields: nil,
			Fields: []database.Field{
				{
					Field:            "t31",
					Style:            "normal",
					Prefix:           "",
					Suffix:           ", ",
					TexValue:         false,
					CitaviAttributes: []string{"author"},
				},
				{
					Field:            "t32",
					Style:            "italic",
					Prefix:           " ",
					Suffix:           ".",
					TexValue:         false,
					CitaviAttributes: nil,
				},
			},
			CiteFields: []database.Field{
				{
					Field:            "t33",
					Style:            "normal",
					Prefix:           "",
					Suffix:           ", ",
					TexValue:         false,
					CitaviAttributes: nil,
				},
			},
		},
	}

	if !reflect.DeepEqual(expected, categories) {
		t.Errorf("expected %v but got %v", expected, categories)
	}
}
