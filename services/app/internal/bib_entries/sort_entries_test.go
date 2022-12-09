package bib_entries

import (
	"reflect"
	"testing"

	"github.com/TimoSto/ThesorTeX/services/app/internal/database"
)

func TestSortEntries(t *testing.T) {
	e := []database.BibEntry{
		{
			Key: "test",
		},
		{
			Key: "atest",
		},
		{
			Key: "btest",
		},
		{
			Key: "aatest",
		},
	}

	es := []database.BibEntry{
		{
			Key: "aatest",
		},
		{
			Key: "atest",
		},
		{
			Key: "btest",
		},
		{
			Key: "test",
		},
	}

	sorted := SortEntries(e)

	if !reflect.DeepEqual(sorted, es) {
		t.Errorf("expected %v to equal %v", sorted, es)
	}
}
