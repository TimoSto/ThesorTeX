package main

import (
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

	mdForVuetify, err := parser.GenerateMdForVuetify(srcFile)
	if err != nil {
		log.Fatalf("could not generate md for vuetify: %v", err)
	}

	f, err := os.Create(filepath.Join(*outDir, "md_for_vuetify.md"))

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err = f.Write(mdForVuetify)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("done")
}
