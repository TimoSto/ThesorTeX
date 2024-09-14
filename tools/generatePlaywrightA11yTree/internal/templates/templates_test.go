package templates

import (
	_ "embed"
	"github.com/google/go-cmp/cmp"
	"testing"
)

//go:embed files/testfile.ts
var testfile string

//go:embed files/testfile_exp.ts
var testfileExp string

func TestGenerateGetTreeTestFile(t *testing.T) {
	got := GenerateGetTreeTestFile(testfile)

	if diff := cmp.Diff(got, testfileExp); diff != "" {
		t.Errorf("unexpected diff: %s", diff)
	}
}
