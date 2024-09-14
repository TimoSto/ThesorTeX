package templates

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed files/getTree.ts
var getTreeTest string

var startString = "test('test', async ({page}) => {"

var startComment = "// Here goes the code from playwright codegen"

func GenerateGetTreeTestFile(generated string) string {
	startIndex := strings.Index(generated, startString) + len(startString)
	endIndex := strings.LastIndex(generated, "\n});")

	testContent := generated[startIndex:endIndex]

	testWithContent := strings.Replace(getTreeTest, startComment, startComment+"\n"+testContent, 1)

	fmt.Println(testWithContent)

	return testWithContent
}
