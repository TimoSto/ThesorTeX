package main

import (
	"flag"
	"strings"

	"github.com/TimoSto/ThesorTeX/tools/documentationcombinator/builder"
)

func main() {
	outDir := flag.String("out-dir", "", "the dir to write the results to")
	docPaths := flag.String("docs", "", "doc paths")
	titlepage := flag.String("title-page", "", "path to the title page")

	flag.Parse()

	docPathsSlice := strings.Split(*docPaths, ",")

	body, err := builder.CombineDocs(docPathsSlice)
	if err != nil {
		panic(err)
	}

	err = builder.BuildDocumentationFromTemplate(*outDir, body, *titlepage)
	if err != nil {
		panic(err)
	}
}
