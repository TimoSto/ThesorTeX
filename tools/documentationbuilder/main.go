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
- MD to be translated into vuetify (see website => frontend => components => DocumentationPanel.vue)
- MD to be used for a reveal.js presentation (see vue-component-library => RevealJS)
- TeX code to be used as content for the template (see project_template)

This go_binary will be used in a bazel genrule
*/

type DocumentationPack struct {
	Title string
	Docs  []parser.DocBody
}

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

	title, rawDocs := parser.SplitDocs(string(srcFile))

	parsedObjects := parser.ParseDocBodies(rawDocs)

	docsPack := DocumentationPack{
		Title: title,
		Docs:  parsedObjects,
	}

	f, err := os.Create(filepath.Join(*outDir, "parsed.json"))

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	data, err := json.MarshalIndent(docsPack, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.Write(data)

	if err != nil {
		log.Fatal(err)
	}

	f2, err := os.Create(filepath.Join(*outDir, "parsed.tex"))

	if err != nil {
		log.Fatal(err)
	}

	defer f2.Close()

	tex, err := parser.GenerateContentForTeX(title, parsedObjects)

	if err != nil {
		log.Fatal(err)
	}

	_, err = f2.Write(tex)

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
