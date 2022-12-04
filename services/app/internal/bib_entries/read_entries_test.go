package bib_entries

import (
	"reflect"
	"testing"

	"github.com/TimoSto/ThesorTeX/mocks/mock_fs"
	"github.com/TimoSto/ThesorTeX/services/app/conf"
)

func TestReadEntriesOfProject_ThreeEntries(t *testing.T) {
	entries, err := ReadEntriesOfProject("test_threeEntries", mock_fs.ReadFile, conf.Config{ProjectsDir: "/test"})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(entries) != 3 {
		t.Errorf("expected three entries, got %v", len(entries))
	}
	expected := []BibEntry{
		{
			Key:        "test1",
			Category:   "testc",
			Fields:     []string{"a", "bb", "ccc"},
			Comment:    "hallo",
			CiteNumber: 0,
		},
		{
			Key:        "test2",
			Category:   "testc",
			Fields:     []string{"a", "bb", "ccc", "ddd"},
			Comment:    "hallo",
			CiteNumber: 0,
		},
		{
			Key:        "test3",
			Category:   "testc2",
			Fields:     []string{"a", "bb", "ccc"},
			Comment:    "",
			CiteNumber: 0,
		},
	}

	if !reflect.DeepEqual(entries, expected) {
		t.Errorf("expected %v, but got %v", expected, entries)
	}
}
