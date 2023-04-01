package parser

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var simpleExpected = []RawDocs{
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
		Content: "Simple\n\ncontent *with* **some** ***styling***",
	},
}

var withCodeExpected = []RawDocs{
	{
		Title:   "Doc 1",
		Content: "Some\ncontent\n```latex\n\\testcommdand{}\nhallo\n```\ntest",
	},
}

func TestSplitDocs(t *testing.T) {
	file, err := os.ReadFile("../testfiles/simple.md")
	if err != nil {
		t.Fatal(err)
	}

	result := SplitDocs(string(file))

	if diff := cmp.Diff(simpleExpected, result); diff != "" {
		t.Errorf("%s", diff)
	}
}

func TestSplitDocsWithCode(t *testing.T) {
	file, err := os.ReadFile("../testfiles/withCode.md")
	if err != nil {
		t.Fatal(err)
	}

	result := SplitDocs(string(file))

	if diff := cmp.Diff(withCodeExpected, result); diff != "" {
		t.Errorf("%s", diff)
	}
}
