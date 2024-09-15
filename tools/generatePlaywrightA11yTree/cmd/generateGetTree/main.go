package main

import (
	"github.com/TimoSto/ThesorTeX/tools/generatePlaywrightA11yTree/internal/templates"
	"os"
)

const generatedFile = "../../tests/playwright/generated/test.spec.ts"
const createdFile = "../../tests/playwright/generated/getA11yTree.spec.ts"

const importString = "import { getAccessibilityTree } from \"../src/a11yTree/a11yTree\";\nimport * as fs from \"node:fs\";"

func main() {
	dat, err := os.ReadFile(generatedFile)
	if err != nil {
		panic(err)
	}

	file := string(dat)

	testFile := templates.GenerateGetTreeTestFile(file)

	// fmt.Println(testFile)

	err = os.WriteFile(createdFile, []byte(testFile), 0666)
	if err != nil {
		panic(err)
	}

}
