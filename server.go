package goat

import (
	"fmt"
	"net/http"
)

type Server struct {
	mux         *http.ServeMux
	controllers []Controller

	DefaultErrorHandler func(int, error, http.ResponseWriter)
}

func NewServer() *Server {
	return &Server{
		mux:         http.NewServeMux(),
		controllers: make([]Controller, 0),
		DefaultErrorHandler: func(i int, err error, w http.ResponseWriter) {
			// TODO: improve
			w.WriteHeader(i)
			fmt.Fprintf(w, "%s", err.Error())
		},
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
