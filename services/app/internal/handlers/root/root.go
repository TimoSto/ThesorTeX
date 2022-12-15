package root

import (
	"fmt"
	"net/http"
)

func HandleRoot() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("root-request", r.URL.Path)
	}

	return http.HandlerFunc(fn)
}
