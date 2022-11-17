package handlers

import (
	"fmt"
	"net/http"
)

func HandleRoot() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("root-request")
	}

	return http.HandlerFunc(fn)
}
