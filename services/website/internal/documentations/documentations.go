package documentations

import (
	"embed"
	"fmt"
)

//go:embed docs
var docs embed.FS

func GetDoc(doc string, lang string) ([]byte, error) {
	return docs.ReadFile(fmt.Sprintf("docs/%s/%s.md", doc, lang))
}
