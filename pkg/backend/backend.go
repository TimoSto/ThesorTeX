package backend

import (
	"github.com/TimoSto/ThesorTeX/pkg/server"
	"github.com/TimoSto/ThesorTeX/pkg/static_content"
)

func Start() {
	srv := server.New("8448") //TODO: port from config

	static_content.Register(srv)

	finished := make(chan bool)
	srv.Start(finished)
}
