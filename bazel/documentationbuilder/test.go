package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	outDir := flag.String("out-dir", "", "the dir to write the results to")
	srcFilePath := flag.String("src", "", "src file")

	flag.Parse()

	fmt.Println(os.Getwd())

	fmt.Println("Reading src file...")

	_, err := os.ReadFile(*srcFilePath)
	if err != nil {
		panic(err)
	}

	fmt.Println(outDir)

}
