package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

/**
Tool to convert a documentation in MD format into
- MD to be translated into vuetify (see vue-component-library => MarkdownToVuetify)
- MD to be used for a reveal.js presentation (see vue-component-library => RevealJS)
- TeX code to be used as content for the template (see project_template)

This go_binary will be used in a bazel genrule
*/

func main() {
	outDir := flag.String("out-dir", "", "out-dir")
	flag.Parse()

	fmt.Println(*outDir)

	fmt.Println(os.Getwd())

	f, err := os.Create(filepath.Join(*outDir, "data.txt"))

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString("old falcon\n")

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("done")
}
