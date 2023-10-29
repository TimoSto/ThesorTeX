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
	title := flag.String("title", "", "title of the document")
	author := flag.String("author", "", "author of the document")
	lang := flag.String("lang", "", "language of the document")
	stripParts := flag.String("strip-parts", "", "parts of the template not to include")

	flag.Parse()

	stripPartsValues := strings.Split(*stripParts, ",")

	docPathsSlice := strings.Split(*docPaths, ",")

	body, err := builder.CombineDocs(docPathsSlice)
	if err != nil {
		panic(err)
	}

	err = builder.BuildDocumentationFromTemplate(*outDir, body, *titlepage, *title, *author, *lang, stripPartsValues)
	if err != nil {
		panic(err)
	}
}
