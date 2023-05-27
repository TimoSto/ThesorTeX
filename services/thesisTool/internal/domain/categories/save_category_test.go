package categories

import (
	"encoding/json"
	"github.com/TimoSto/ThesorTeX/services/thesisTool/internal/domain/projects"
	"reflect"
	"testing"

	"github.com/TimoSto/ThesorTeX/pkg/backend/filesystem/fake"
	"github.com/TimoSto/ThesorTeX/pkg/backend/pathbuilder"
	"github.com/TimoSto/ThesorTeX/services/thesisTool/internal/config"
)

func TestSaveCategory_Override(t *testing.T) {
	fs := fake.FileSystem{}

	cfg := config.Config{ProjectsDir: "/projects"}

	file := "[\n  {\n    \"Name\": \"aufsatz\",\n    \"CitaviCategory\": \"\",\n    \"CitaviFilters\": null,\n    \"BibFields\": [\n      {\n        \"Name\": \"Autor\",\n        \"Format\": {\n          \"Italic\": true,\n          \"Prefix\": \"\",\n          \"Suffix\": \" \",\n          \"Preformatted\": false\n        },\n        \"CitaviMapping\": null\n      },\n      {\n        \"Name\": \"Jahr\",\n        \"Format\": {\n          \"Italic\": false,\n          \"Prefix\": \"(\",\n          \"Suffix\": \"), \",\n          \"Preformatted\": false\n        },\n        \"CitaviMapping\": null\n      },\n      {\n        \"Name\": \"Titel\",\n        \"Format\": {\n          \"Italic\": false,\n          \"Prefix\": \"\",\n          \"Suffix\": \", \",\n          \"Preformatted\": false\n        },\n        \"CitaviMapping\": null\n      },\n      {\n        \"Name\": \"Zeitschrift\",\n        \"Format\": {\n          \"Italic\": true,\n          \"Prefix\": \"in: \",\n          \"Suffix\": \" \",\n          \"Preformatted\": false\n        },\n        \"CitaviMapping\": null\n      },\n      {\n        \"Name\": \"Ausgabe\",\n        \"Format\": {\n          \"Italic\": false,\n          \"Prefix\": \"\",\n          \"Suffix\": \" \",\n          \"Preformatted\": false\n        },\n        \"CitaviMapping\": null\n      },\n      {\n        \"Name\": \"Seiten\",\n        \"Format\": {\n          \"Italic\": true,\n          \"Prefix\": \"S. \",\n          \"Suffix\": \", \",\n          \"Preformatted\": false\n        },\n        \"CitaviMapping\": null\n      },\n      {\n        \"Name\": \"url\",\n        \"Format\": {\n          \"Italic\": true,\n          \"Prefix\": \"abrufbar unter: \",\n          \"Suffix\": \" \",\n          \"Preformatted\": false\n        },\n        \"CitaviMapping\": null\n      }\n    ],\n    \"CiteFields\": [\n      {\n        \"Name\": \"Autor_kurz\",\n        \"Format\": {\n          \"Italic\": true,\n          \"Prefix\": \"\",\n          \"Suffix\": \" \",\n          \"Preformatted\": false\n        },\n        \"CitaviMapping\": null\n      },\n      {\n        \"Name\": \"Jahr\",\n        \"Format\": {\n          \"Italic\": true,\n          \"Prefix\": \"\",\n          \"Suffix\": \" \",\n          \"Preformatted\": false\n        },\n        \"CitaviMapping\": null\n      }\n    ]\n  }\n]"

	fs.WriteFile(pathbuilder.GetPathInProject(cfg.ProjectsDir, "test", categoriesFile), []byte(file))

	projects.CreateFakeMetaFile(cfg, fs, "test")

	c := Category{
		Name:           "aufsatz",
		BibFields:      []field{},
		CiteFields:     []field{},
		CitaviCategory: "",
		CitaviFilter:   nil,
	}

	err := SaveCategory(&fs, cfg, "test", "aufsatz", c)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	resultFile, _ := fs.ReadFile(pathbuilder.GetPathInProject(cfg.ProjectsDir, "test", categoriesFile))

	var result []Category
	json.Unmarshal(resultFile, &result)

	if !reflect.DeepEqual(result, []Category{c}) {
		t.Errorf("expected to only find category %v, but found %v", c, result)
	}
}

func TestSaveCategory_Rename(t *testing.T) {
	fs := fake.FileSystem{}

	cfg := config.Config{ProjectsDir: "/projects"}

	file := "[\n  {\n    \"Name\": \"aufsatz\",\n    \"CitaviCategory\": \"\",\n    \"CitaviFilters\": null,\n    \"BibFields\": [\n      {\n        \"Name\": \"Autor\",\n        \"Format\": {\n          \"Italic\": true,\n          \"Prefix\": \"\",\n          \"Suffix\": \" \",\n          \"Preformatted\": false\n        },\n        \"CitaviMapping\": null\n      },\n      {\n        \"Name\": \"Jahr\",\n        \"Format\": {\n          \"Italic\": false,\n          \"Prefix\": \"(\",\n          \"Suffix\": \"), \",\n          \"Preformatted\": false\n        },\n        \"CitaviMapping\": null\n      },\n      {\n        \"Name\": \"Titel\",\n        \"Format\": {\n          \"Italic\": false,\n          \"Prefix\": \"\",\n          \"Suffix\": \", \",\n          \"Preformatted\": false\n        },\n        \"CitaviMapping\": null\n      },\n      {\n        \"Name\": \"Zeitschrift\",\n        \"Format\": {\n          \"Italic\": true,\n          \"Prefix\": \"in: \",\n          \"Suffix\": \" \",\n          \"Preformatted\": false\n        },\n        \"CitaviMapping\": null\n      },\n      {\n        \"Name\": \"Ausgabe\",\n        \"Format\": {\n          \"Italic\": false,\n          \"Prefix\": \"\",\n          \"Suffix\": \" \",\n          \"Preformatted\": false\n        },\n        \"CitaviMapping\": null\n      },\n      {\n        \"Name\": \"Seiten\",\n        \"Format\": {\n          \"Italic\": true,\n          \"Prefix\": \"S. \",\n          \"Suffix\": \", \",\n          \"Preformatted\": false\n        },\n        \"CitaviMapping\": null\n      },\n      {\n        \"Name\": \"url\",\n        \"Format\": {\n          \"Italic\": true,\n          \"Prefix\": \"abrufbar unter: \",\n          \"Suffix\": \" \",\n          \"Preformatted\": false\n        },\n        \"CitaviMapping\": null\n      }\n    ],\n    \"CiteFields\": [\n      {\n        \"Name\": \"Autor_kurz\",\n        \"Format\": {\n          \"Italic\": true,\n          \"Prefix\": \"\",\n          \"Suffix\": \" \",\n          \"Preformatted\": false\n        },\n        \"CitaviMapping\": null\n      },\n      {\n        \"Name\": \"Jahr\",\n        \"Format\": {\n          \"Italic\": true,\n          \"Prefix\": \"\",\n          \"Suffix\": \" \",\n          \"Preformatted\": false\n        },\n        \"CitaviMapping\": null\n      }\n    ]\n  }\n]"

	fs.WriteFile(pathbuilder.GetPathInProject(cfg.ProjectsDir, "test", categoriesFile), []byte(file))

	projects.CreateFakeMetaFile(cfg, fs, "test")

	c := Category{
		Name:           "aufsatzf",
		BibFields:      []field{},
		CiteFields:     []field{},
		CitaviCategory: "",
		CitaviFilter:   nil,
	}

	err := SaveCategory(&fs, cfg, "test", "aufsatz", c)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	resultFile, _ := fs.ReadFile(pathbuilder.GetPathInProject(cfg.ProjectsDir, "test", categoriesFile))

	var result []Category
	json.Unmarshal(resultFile, &result)

	if !reflect.DeepEqual(result, []Category{c}) {
		t.Errorf("expected to only find category %v, but found %v", c, result)
	}
}

func TestSaveCategory_Add(t *testing.T) {
	fs := fake.FileSystem{}

	cfg := config.Config{ProjectsDir: "/projects"}

	file := "[\n  {\n    \"Name\": \"aufsatz\",\n    \"CitaviCategory\": \"\",\n    \"CitaviFilters\": null,\n    \"BibFields\": [\n      {\n        \"Name\": \"Autor\",\n        \"Format\": {\n          \"Italic\": true,\n          \"Prefix\": \"\",\n          \"Suffix\": \" \",\n          \"Preformatted\": false\n        },\n        \"CitaviMapping\": null\n      },\n      {\n        \"Name\": \"Jahr\",\n        \"Format\": {\n          \"Italic\": false,\n          \"Prefix\": \"(\",\n          \"Suffix\": \"), \",\n          \"Preformatted\": false\n        },\n        \"CitaviMapping\": null\n      },\n      {\n        \"Name\": \"Titel\",\n        \"Format\": {\n          \"Italic\": false,\n          \"Prefix\": \"\",\n          \"Suffix\": \", \",\n          \"Preformatted\": false\n        },\n        \"CitaviMapping\": null\n      },\n      {\n        \"Name\": \"Zeitschrift\",\n        \"Format\": {\n          \"Italic\": true,\n          \"Prefix\": \"in: \",\n          \"Suffix\": \" \",\n          \"Preformatted\": false\n        },\n        \"CitaviMapping\": null\n      },\n      {\n        \"Name\": \"Ausgabe\",\n        \"Format\": {\n          \"Italic\": false,\n          \"Prefix\": \"\",\n          \"Suffix\": \" \",\n          \"Preformatted\": false\n        },\n        \"CitaviMapping\": null\n      },\n      {\n        \"Name\": \"Seiten\",\n        \"Format\": {\n          \"Italic\": true,\n          \"Prefix\": \"S. \",\n          \"Suffix\": \", \",\n          \"Preformatted\": false\n        },\n        \"CitaviMapping\": null\n      },\n      {\n        \"Name\": \"url\",\n        \"Format\": {\n          \"Italic\": true,\n          \"Prefix\": \"abrufbar unter: \",\n          \"Suffix\": \" \",\n          \"Preformatted\": false\n        },\n        \"CitaviMapping\": null\n      }\n    ],\n    \"CiteFields\": [\n      {\n        \"Name\": \"Autor_kurz\",\n        \"Format\": {\n          \"Italic\": true,\n          \"Prefix\": \"\",\n          \"Suffix\": \" \",\n          \"Preformatted\": false\n        },\n        \"CitaviMapping\": null\n      },\n      {\n        \"Name\": \"Jahr\",\n        \"Format\": {\n          \"Italic\": true,\n          \"Prefix\": \"\",\n          \"Suffix\": \" \",\n          \"Preformatted\": false\n        },\n        \"CitaviMapping\": null\n      }\n    ]\n  }\n]"

	fs.WriteFile(pathbuilder.GetPathInProject(cfg.ProjectsDir, "test", categoriesFile), []byte(file))

	projects.CreateFakeMetaFile(cfg, fs, "test")

	c := Category{
		Name:           "aufsatzf",
		BibFields:      []field{},
		CiteFields:     []field{},
		CitaviCategory: "",
		CitaviFilter:   nil,
	}

	err := SaveCategory(&fs, cfg, "test", "", c)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	resultFile, _ := fs.ReadFile(pathbuilder.GetPathInProject(cfg.ProjectsDir, "test", categoriesFile))

	var result []Category
	json.Unmarshal(resultFile, &result)

	if len(result) != 2 {
		t.Errorf("expected two categories, got %v", len(result))
	}

	if !reflect.DeepEqual(result[1], c) {
		t.Errorf("expected to find category %v, but found %v", c, result[1])
	}
}
