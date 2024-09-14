package main

import (
	"github.com/TimoSto/ThesorTeX/tools/generatePlaywrightA11yTree/internal/templates"
	"os"
)

const generatedFile = "../../tests/playwright/generated/test.spec.ts"
const createdFile = "../../tests/playwright/generated/assertA11yTree.spec.ts"

func main() {
	dat, err := os.ReadFile(generatedFile)
	if err != nil {
		panic(err)
	}

	file := string(dat)

	testFile := templates.GenerateAssertTreeTestFile(file)

	// fmt.Println(testFile)

	err = os.WriteFile(createdFile, []byte(testFile), 0666)
	if err != nil {
		panic(err)
	}

}
