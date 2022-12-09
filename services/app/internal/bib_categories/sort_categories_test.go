package bib_categories

import (
	"reflect"
	"testing"

	"github.com/TimoSto/ThesorTeX/services/app/internal/database"
)

func TestSortCategories(t *testing.T) {
	e := []database.BibCategory{
		{
			Name: "test",
		},
		{
			Name: "atest",
		},
		{
			Name: "btest",
		},
		{
			Name: "aatest",
		},
	}

	es := []database.BibCategory{
		{
			Name: "aatest",
		},
		{
			Name: "atest",
		},
		{
			Name: "btest",
		},
		{
			Name: "test",
		},
	}

	sorted := SortCategories(e)

	if !reflect.DeepEqual(sorted, es) {
		t.Errorf("expected %v to equal %v", sorted, es)
	}
}
