package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello world", os.Args)
	args := os.Args[1:]

	err := os.MkdirAll(args[0], os.ModePerm)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(args[0]+"/test.txt", []byte("hello"), os.ModePerm)
	if err != nil {
		panic(err)
	}

}
