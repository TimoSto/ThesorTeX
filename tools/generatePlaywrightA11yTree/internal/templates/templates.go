package templates

import (
	_ "embed"
	"strings"
)

//go:embed files/getTree.ts
var getTreeTest string

//go:embed files/assertA11yTree.spec.ts
var assertTreeTest string

var startComment = "// Here goes the code from playwright codegen"

func GenerateGetTreeTestFile(generated string) string {

	lines := strings.Split(generated, "\n")

	lines = lines[3 : len(lines)-1]

	testContent := strings.Join(lines, "\n")

	testWithContent := strings.Replace(getTreeTest, startComment, startComment+"\n"+testContent, 1)

	return testWithContent
}

func GenerateAssertTreeTestFile(generated string) string {

	lines := strings.Split(generated, "\n")

	lines = lines[3 : len(lines)-1]

	testContent := strings.Join(lines, "\n")

	testWithContent := strings.Replace(assertTreeTest, startComment, startComment+"\n"+testContent, 1)

	return testWithContent
}
