package main

import (
	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"os/exec"
)

func main() {
	_, err := exec.LookPath("npx")
	if err != nil {
		log.Fatal("could not find npx installation")
	}

	codegen := exec.Command("npx", "playwright", "codegen")
	err = codegen.Run()
	if err != nil {
		log.Error("Make sure npx and playwright are installed globally")
		log.Fatal("could not start codegen: ", err)
	}
}
