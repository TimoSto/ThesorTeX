package categories

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/TimoSto/ThesorTeX/pkg/backend/filesystem/fake"
	"github.com/TimoSto/ThesorTeX/pkg/backend/pathbuilder"
	"github.com/TimoSto/ThesorTeX/services/thesisTool/internal/config"
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
						Italic:       false,
						Preformatted: false,
					},
				},
				{
					Name:          "f2",
					CitaviMapping: []string{},
					Format: format{
						Prefix:       "(",
						Suffix:       ")",
						Italic:       true,
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
						Italic:       false,
						Preformatted: false,
					},
				},
				{
					Name:          "f3",
					CitaviMapping: []string{},
					Format: format{
						Prefix:       "- ",
						Suffix:       "",
						Italic:       true,
						Preformatted: false,
					},
				},
			},
		},
	}
	data, _ := json.Marshal(existing)

	fs.WriteFile(pathbuilder.GetPathInProject(cfg.ProjectsDir, "test", categoriesFile), data)

	got, err := GetAllCategories("test", &fs, cfg)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !reflect.DeepEqual(got, existing) {
		t.Errorf("expected %v, got %v", existing, got)
	}

}
