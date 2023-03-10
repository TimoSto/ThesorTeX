package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/TimoSto/ThesorTeX/tools/releasenotes/parser"
)

var target = flag.String("target", "", "TOOL, TEMPLATE, CV_TEMPLATE")

func init() {
	flag.Parse()

	if *target == "" {
		log.Fatal("Flag \"-target\" not set.")
	}
}

func main() {
	versionsFile, err := os.ReadFile("VERSIONS")
	if err != nil {
		log.Panicf("could not read versions file: %v", err)
	}

	v := parser.GetVersions(versionsFile, *target)

	file, err := os.ReadFile(fmt.Sprintf("CHANGELOG_%s.md", *target))
	if err != nil {
		log.Panicf("could not read changelog file: %v", err)
	}

	notes := parser.GetRelevantNotes(file, v)

	fmt.Println(notes)
}
