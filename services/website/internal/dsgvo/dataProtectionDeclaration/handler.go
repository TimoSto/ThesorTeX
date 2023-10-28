package dataProtectionDeclaration

import (
	_ "embed"
	"net/http"
)

//go:embed de/main.pdf
var declarationDe string

func Register(mux *http.ServeMux) {
	mux.HandleFunc("/dsgvo", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(declarationDe))
	})
}
