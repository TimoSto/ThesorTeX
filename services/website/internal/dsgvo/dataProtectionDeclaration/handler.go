package dataProtectionDeclaration

import (
	_ "embed"
	"net/http"
)

//go:embed dpd_de.pdf
var declarationDePDF []byte

//go:embed dpd/de/parsed.json
var declarationDeJSON []byte

func HandleDSGVO() func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		switch r.Header.Get("Accept") {
		case "application/json":
			w.Write(declarationDeJSON)
		default:
			w.Write(declarationDePDF)
		}
	}

	return fn
}
