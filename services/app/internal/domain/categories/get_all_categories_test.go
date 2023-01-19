package categories

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/TimoSto/ThesorTeX/pkg/backend/pathbuilder"
	"github.com/TimoSto/ThesorTeX/services/app/internal/config"
	"github.com/TimoSto/ThesorTeX/services/app/internal/filesystem/fake"
)

func TestGetAllCategories(t *testing.T) {
	fs := fake.FileSystem{}

	cfg := config.Config{
		ProjectsDir: "projects/",
	}

	existing := []Category{
		{
			Name:           "c1",
			CitaviCategory: "",
			BibFields: []field{
				{
					Name:          "f1",
					CitaviMapping: []string{},
					Format: format{
						Prefix:       "",
						Suffix:       " ",
						Style:        "normal",
						Preformatted: false,
					},
				},
				{
					Name:          "f2",
					CitaviMapping: []string{},
					Format: format{
						Prefix:       "(",
						Suffix:       ")",
						Style:        "italic",
						Preformatted: false,
					},
				},
			},
			CiteFields: []field{
				{
					Name:          "f1",
					CitaviMapping: []string{},
					Format: format{
						Prefix:       "",
						Suffix:       " ",
						Style:        "normal",
						Preformatted: false,
					},
				},
				{
					Name:          "f3",
					CitaviMapping: []string{},
					Format: format{
						Prefix:       "- ",
						Suffix:       "",
						Style:        "italic",
						Preformatted: false,
					},
				},
			},
		},
	}
	data, _ := json.Marshal(existing)

	fs.WriteFile(pathbuilder.GetPathInProject(cfg.ProjectsDir, "test", categoriesFile), data)

	got, err := getAllCategories("test", &fs, cfg)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !reflect.DeepEqual(got, existing) {
		t.Errorf("expected %v, got %v", existing, got)
	}

}
