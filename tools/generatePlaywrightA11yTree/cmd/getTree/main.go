package main

import (
	"fmt"
	"os"
	"strings"
)

const generatedFile = "../../tests/playwright/generated/test.spec.ts"

const importString = "import { getAccessibilityTree } from \"../src/a11yTree/a11yTree\";\nimport * as fs from \"node:fs\";"

func main() {
	dat, err := os.ReadFile(generatedFile)
	if err != nil {
		panic(err)
	}

	fmt.Println("got test file:")
	file := string(dat)

	fmt.Println(file)

	fmt.Println("adding import...")
	file = importString + file

	fmt.Println("adding get of tree...")

	endIndex := strings.LastIndex(file, "});")

	file = file[:endIndex] +
		"\n\tconst client = await page.context().newCDPSession(page);\n\n\tconst tree = await getAccessibilityTree(client);\n" +
		"\n\tconst file = JSON.stringify(tree, null, 2);" +
		"\tfs.writeFile(\"./generated/a11yTree.json\", file, (err) => {\n\t\texpect(err).toBeNull();\n\t});" +
		file[endIndex:]

	err = os.WriteFile(generatedFile, []byte(file), 0666)
	if err != nil {
		panic(err)
	}

}
