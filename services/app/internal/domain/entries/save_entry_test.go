package entries

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/TimoSto/ThesorTeX/pkg/backend/pathbuilder"
	"github.com/TimoSto/ThesorTeX/services/app/internal/config"
	"github.com/TimoSto/ThesorTeX/services/app/internal/filesystem/fake"
)

func TestSaveEntry_Override(t *testing.T) {
	fs := fake.FileSystem{}

	cfg := config.Config{ProjectsDir: "/projects"}

	file := "[\n  {\n    \"Key\": \"testEntry\",\n    \"Category\": \"aufsatz\",\n    \"Fields\": [\n      \"Autor xyz\",\n      \"1999\",\n      \"ThesorTeX - Ein tolles Tool\",\n      \"Random Zeitschrift\",\n      \"99/03\",\n      \"20-22\",\n      \"https://wikipedia.org/\"\n    ]\n  }\n]"

	fs.WriteFile(pathbuilder.GetPathInProject(cfg.ProjectsDir, "test", entriesFile), []byte(file))

	c := Entry{
		Key:      "testEntry",
		Category: "aufsatz",
		Fields:   []string{},
	}

	err := SaveEntry(&fs, cfg, "test", "testEntry", c)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	resultFile, _ := fs.ReadFile(pathbuilder.GetPathInProject(cfg.ProjectsDir, "test", entriesFile))

	var result []Entry
	json.Unmarshal(resultFile, &result)

	if !reflect.DeepEqual(result, []Entry{c}) {
		t.Errorf("expected to only find category %v, but found %v", c, result)
	}
}

func TestSaveCategory_Rename(t *testing.T) {
	fs := fake.FileSystem{}

	cfg := config.Config{ProjectsDir: "/projects"}

	file := "[\n  {\n    \"Key\": \"testEntry\",\n    \"Category\": \"aufsatz\",\n    \"Fields\": [\n      \"Autor xyz\",\n      \"1999\",\n      \"ThesorTeX - Ein tolles Tool\",\n      \"Random Zeitschrift\",\n      \"99/03\",\n      \"20-22\",\n      \"https://wikipedia.org/\"\n    ]\n  }\n]"

	fs.WriteFile(pathbuilder.GetPathInProject(cfg.ProjectsDir, "test", entriesFile), []byte(file))

	c := Entry{
		Key:      "testEntry2",
		Category: "aufsatz",
		Fields:   []string{},
	}

	err := SaveEntry(&fs, cfg, "test", "testEntry", c)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	resultFile, _ := fs.ReadFile(pathbuilder.GetPathInProject(cfg.ProjectsDir, "test", entriesFile))

	var result []Entry
	json.Unmarshal(resultFile, &result)

	if !reflect.DeepEqual(result, []Entry{c}) {
		t.Errorf("expected to only find category %v, but found %v", c, result)
	}
}

func TestSaveCategory_Add(t *testing.T) {
	fs := fake.FileSystem{}

	cfg := config.Config{ProjectsDir: "/projects"}

	file := "[\n  {\n    \"Key\": \"testEntry\",\n    \"Category\": \"aufsatz\",\n    \"Fields\": [\n      \"Autor xyz\",\n      \"1999\",\n      \"ThesorTeX - Ein tolles Tool\",\n      \"Random Zeitschrift\",\n      \"99/03\",\n      \"20-22\",\n      \"https://wikipedia.org/\"\n    ]\n  }\n]"

	fs.WriteFile(pathbuilder.GetPathInProject(cfg.ProjectsDir, "test", entriesFile), []byte(file))

	c := Entry{
		Key:      "testEntryf",
		Category: "aufsatz",
		Fields:   []string{},
	}

	err := SaveEntry(&fs, cfg, "test", "", c)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	resultFile, _ := fs.ReadFile(pathbuilder.GetPathInProject(cfg.ProjectsDir, "test", entriesFile))

	var result []Entry
	json.Unmarshal(resultFile, &result)

	if len(result) != 2 {
		t.Errorf("expected 2 entries, got %v", len(result))
	}

	if !reflect.DeepEqual(result[1], c) {
		t.Errorf("expected to only find category %v, but found %v", c, result)
	}
}