package backend

import (
	"fmt"
	"net/http"

	"github.com/TimoSto/ThesorTeX/pkg/server"
)

func Start() {
	srv := server.New("8448") //TODO: port from config

	srv.Register("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hi")
	})

	finished := make(chan bool)
	srv.Start(finished)
}
