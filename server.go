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

func (s *Server) AddController(c Controller) {
	s.controllers = append(s.controllers, c)

	path, method := c.GetPathAndMethod()
	s.mux.Handle(fmt.Sprintf("%s %s", method, path), c.MakeHandlerFunc(s))
}
