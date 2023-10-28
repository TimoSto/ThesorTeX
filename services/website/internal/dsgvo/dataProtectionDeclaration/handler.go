package dataProtectionDeclaration

import (
	_ "embed"
	"fmt"
	"net/http"
)

//go:embed de/main.pdf
var declarationDe string

func HandleDSGVO() func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("got here")
		w.Write([]byte(declarationDe))
	}

	return fn
}
