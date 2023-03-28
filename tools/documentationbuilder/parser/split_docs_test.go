package parser

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSplitDocs(t *testing.T) {
	file, err := os.ReadFile("../testfiles/simple.md")
	if err != nil {
		t.Fatal(err)
	}

	result := SplitDocs(string(file))

	expected := []RawDocs{
		{
			Title:   "Doc 1",
			Content: "Some\ncontent",
		},
		{
			Title:   "Doc 2 styled",
			Content: "Simple\ncontent *with* **some** ***styling***",
		},
		{
			Title:   "Extra spacing",
			Content: "Simple\ncontent *with* **some** ***styling***",
		},
	}

	if diff := cmp.Diff(expected, result); diff != "" {
		t.Errorf("%s", diff)
	}
}
