package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/TimoSto/ThesorTeX/tools/documentationbuilder/parser"
)

/**
Tool to convert a documentation in MD format into
- MD to be translated into vuetify (see vue-component-library => MarkdownToVuetify)
- MD to be used for a reveal.js presentation (see vue-component-library => RevealJS)
- TeX code to be used as content for the template (see project_template)

This go_binary will be used in a bazel genrule
*/

func main() {
	outDir := flag.String("out-dir", "", "the dir to write the results to")
	srcFilePath := flag.String("src", "", "src file")

	flag.Parse()

	fmt.Println("Reading src file...")

	srcFile, err := os.ReadFile(*srcFilePath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("converting md to md for vuetify...")

	rawDocs := parser.SplitDocs(string(srcFile))

	parsedObjects := parser.ParseDocBodies(rawDocs)

	f, err := os.Create(filepath.Join(*outDir, "parsed.json"))

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	data, err := json.MarshalIndent(parsedObjects, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.Write(data)

	if err != nil {
		log.Fatal(err)
	}

	err = parser.BuildDocumentationFromTemplate(*outDir)

	if err != nil {
		log.Fatal(err)
	}

	//mdForVuetify, err := parser.GenerateMdForVuetify(srcFile)
	//if err != nil {
	//	log.Fatalf("could not generate md for vuetify: %v", err)
	//}
	//

	//
	//if err != nil {
	//	log.Fatal(err)
	//}

	fmt.Println("done")
}
