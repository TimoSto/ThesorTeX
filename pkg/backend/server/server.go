package server

//TODO use this package or delete it

import (
	"fmt"
	"net"
	"net/http"

	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
)

type Server struct {
	handler http.Handler
	port    string
	srv     *http.Server
}

func New(port string, chain http.Handler) Server {
	return Server{
		handler: chain,
		port:    port,
	}
}

func (s *Server) Start(fin chan bool) string {
	s.srv = &http.Server{
		Handler: s.handler,
	}

	socket, err := net.Listen("tcp", "localhost:"+s.port)
	if err != nil {
		panic(fmt.Sprintf("server: could not bind to socket (port: %v)", s.port))
	}

	go func() {
		err := s.srv.Serve(socket)
		if err != nil {
			if err != http.ErrServerClosed {
				log.Error("Error starting server: %v", err)
			}
			fin <- true
		}
	}()

	_, port, _ := net.SplitHostPort(socket.Addr().String())

	log.Info(fmt.Sprintf("Server running at http://localhost:%v", port))

	return fmt.Sprintf("http://localhost:%v", port)
}
