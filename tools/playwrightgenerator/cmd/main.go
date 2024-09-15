package main

import (
	"fmt"
	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"os"
	"os/exec"
	"path"
)

func main() {
	_, err := exec.LookPath("npx")
	if err != nil {
		log.Fatal("could not find npx installation")
	}

	wd, _ := os.Getwd()
	outPath := path.Join(wd, "generated/tests.spec.ts")

	err = os.MkdirAll(path.Join(wd, "generated"), os.ModePerm)
	if err != nil {
		log.Fatal("could not create dir: ", err)
	}

	_, err = os.OpenFile(outPath, os.O_RDONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal("could not create file: ", err)
	}

	fmt.Println(outPath)

	codegen := exec.Command("npx", "playwright", "codegen", "-o", outPath)
	err = codegen.Run()
	if err != nil {
		log.Info("Make sure playwright and its browsers are installed globally")
		log.Fatal("could not start codegen: ", err)
	}
}
