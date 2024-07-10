package goat

import (
	"fmt"
	"net/http"
	"reflect"
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

func (Route[Params, Return]) GenerateSwagger() ([]byte, error) {
	panic("not implemented")
}

func (route Route[Params, Return]) MakeHandlerFunc(s *Server) http.HandlerFunc {
	var sampleParams Params
	route.blueprints = compileBlueprints(sampleParams)
	// fmt.Printf("%#v", route.blueprints)

	return func(w http.ResponseWriter, r *http.Request) {
		params := reflect.ValueOf(&sampleParams).Elem()

		for _, b := range route.blueprints {
			err := b.SetField(params, s, r)
			if err != nil {
				s.DefaultErrorHandler(400, err, w)
				return
			}
		}

		// TEMP
		fmt.Printf("%#v\n", sampleParams)
		return
		// !TEMP

		ctx := Context[Params]{Params: sampleParams}
		status, v, err := route.Handler(&ctx)
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
