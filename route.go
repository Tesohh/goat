package goat

import (
	"fmt"
	"net/http"
)

type Controller interface {
	MakeHandlerFunc(*Server) http.HandlerFunc
	GetPathAndMethod() (string, string)
	GenerateSwagger() ([]byte, error)
}

type Route[Params any, Return any] struct {
	Path               string
	Method             string
	Description        string
	ParamsDescriptions map[string]string
	Handler            func(c *Context[Params]) (status int, v *Return, err error)

	blueprints []fieldBlueprint
}

func (route Route[Params, Return]) GetPathAndMethod() (string, string) {
	return route.Path, route.Method
}

func (route Route[Params, Return]) MakeHandlerFunc(s *Server) http.HandlerFunc {
	var sampleParams Params
	route.blueprints = compileBlueprints(sampleParams)
	fmt.Printf("%#v", route.blueprints)

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := &Context[Params]{}

		status, v, err := route.Handler(ctx)
		if err != nil {
			s.DefaultErrorHandler(status, err, w)
			return
		}

		if v == nil {
			return
		}
	}
}

func GenerateSwagger() ([]byte, error) {
	panic("not implemented")
}
