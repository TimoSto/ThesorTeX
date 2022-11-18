package server

import (
	"fmt"
	"net"
	"net/http"

	"github.com/TimoSto/ThesorTeX/pkg/log"
)

type Server struct {
	mux  *http.ServeMux
	port string
	srv  *http.Server
}

func New(port string) Server {
	return Server{
		mux:  http.NewServeMux(),
		port: port,
	}
}

func (s *Server) Register(path string, handler http.HandlerFunc) {
	s.mux.HandleFunc(path, handler)
}

func (s *Server) Start(fin chan bool) {
	s.srv = &http.Server{
		Handler: s.mux,
	}

	socket, err := net.Listen("tcp", "0.0.0.0:"+s.port)
	if err != nil {
		panic(fmt.Sprintf("server: could not bind to socket (port: %v)", s.port))
	}

	go func() {
		err := s.srv.Serve(socket)
		if err != nil {
			if err != http.ErrServerClosed {
				log.Error("Error starting server", err)
			}
			fin <- true
		}
	}()

	_, port, _ := net.SplitHostPort(socket.Addr().String())

	log.Info(fmt.Sprintf("Server running at http://localhost:%v", port))
}
