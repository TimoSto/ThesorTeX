package server

import (
	"errors"
	"fmt"
	"github.com/TimoSto/ThesorTeX/pkg/backend/log"
	"net"
	"net/http"
)

type Server struct {
	srv  *http.Server
	addr string
}

func NewServer(port string, handler http.Handler, fin chan bool) *Server {
	s := &Server{
		srv: &http.Server{
			Handler: handler,
		},
		addr: "",
	}

	socket, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		panic(fmt.Sprintf("server: could not bind to socket (port: %v)", port))
	}

	go func() {
		fmt.Println("serving")
		err := s.srv.Serve(socket)
		if err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				log.Error("Error starting server: %v", err)
			}
			fin <- true
		}
	}()

	s.addr = socket.Addr().String()

	return s
}

func (s *Server) Address() string {
	return s.addr
}

func (s *Server) Port() string {
	_, port, _ := net.SplitHostPort(s.addr)

	return port
}
