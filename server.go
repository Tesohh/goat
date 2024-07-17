package goat

import (
	"fmt"
	"net/http"
)

type ErrorHandlerFunc func(http.ResponseWriter, int, error)

type EncoderFunc func(http.ResponseWriter, any) // the any part will receive a value, not a pointer

type Server struct {
	mux         *http.ServeMux
	controllers []Controller

	ErrorHandler ErrorHandlerFunc
	Encoder      EncoderFunc
	Info         ServerInfo
}

// Returns a new pointer to a server, with the goat.DefaultErrorHandler and a goat.JSONEncoder.
func NewServer(info ServerInfo) *Server {
	return &Server{
		mux:          http.NewServeMux(),
		controllers:  make([]Controller, 0),
		ErrorHandler: DefaultErrorHandler,
		Encoder:      JSONEncoder,
		Info:         info,
	}
}

func (s *Server) AddController(c Controller) {
	s.controllers = append(s.controllers, c)

	path, method := c.GetPathAndMethod()
	s.mux.Handle(fmt.Sprintf("%s %s", method, path), c.MakeHandlerFunc(s))
}

func (s *Server) Listen(addr string) {
	http.ListenAndServe(addr, s.mux)
}
