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

var withImageExpected = []RawDocs{
	{
		Title:   "Doc 1",
		Content: "Some\ncontent\n![test image](./images/img1.png)\ntest",
	},
}

func TestSplitDocs(t *testing.T) {
	file, err := os.ReadFile("../testfiles/simple.md")
	if err != nil {
		t.Fatal(err)
	}

	title, result := SplitDocs(string(file))

	expTitle := "Simple Docs"

	if title != expTitle {
		t.Errorf("expected title %s, but got %s", expTitle, title)
	}

	if diff := cmp.Diff(simpleExpected, result); diff != "" {
		t.Errorf("%s", diff)
	}
}

func TestSplitDocsWithCode(t *testing.T) {
	file, err := os.ReadFile("../testfiles/withCode.md")
	if err != nil {
		t.Fatal(err)
	}

	title, result := SplitDocs(string(file))

	expTitle := "With code"

	if title != expTitle {
		t.Errorf("expected title %s, but got %s", expTitle, title)
	}

	if diff := cmp.Diff(withCodeExpected, result); diff != "" {
		t.Errorf("%s", diff)
	}
}

func TestSplitDocsWithImage(t *testing.T) {
	file, err := os.ReadFile("../testfiles/withImage.md")
	if err != nil {
		t.Fatal(err)
	}

	title, result := SplitDocs(string(file))

	expTitle := "With image"

	if title != expTitle {
		t.Errorf("expected title %s, but got %s", expTitle, title)
	}

	if diff := cmp.Diff(withImageExpected, result); diff != "" {
		t.Errorf("%s", diff)
	}
}
