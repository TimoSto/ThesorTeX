package templates

import (
	_ "embed"
	"strings"
)

//go:embed files/getTree.ts
var getTreeTest string

var startString = "import { test, expect } from '@playwright/test';\n\ntest('test', async ({ page }) => {"

var startComment = "// Here goes the code from playwright codegen"

func GenerateGetTreeTestFile(generated string) string {

	lines := strings.Split(generated, "\n")

	lines = lines[3 : len(lines)-1]

	testContent := strings.Join(lines, "\n")

	testWithContent := strings.Replace(getTreeTest, startComment, startComment+"\n"+testContent, 1)

	return testWithContent
}
