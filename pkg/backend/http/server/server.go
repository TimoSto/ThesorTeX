package server

import (
	"errors"
	"fmt"
	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"net"
	"net/http"
)

func StartServer(port string, handler http.Handler) (string, error) {
	srv := &http.Server{
		Handler: handler,
	}

	socket, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		return "", fmt.Errorf("server: could not bind to socket (port: %v)", port)
	}

	go func() {
		err := srv.Serve(socket)
		if err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				log.Error("Error starting server: %v", err)
			}
		}
	}()

	_, p, err := net.SplitHostPort(socket.Addr().String())
	if err != nil {
		return "", err
	}

	return p, nil
}
