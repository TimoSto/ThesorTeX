package bib_categories

import (
	"reflect"
	"testing"
)

func TestSortCategories(t *testing.T) {
	e := []BibCategory{
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

	es := []BibCategory{
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
