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

func (route Route[Params, Return]) MakeHandlerFunc(s *Server) http.HandlerFunc {
	var sampleParams Params
	route.blueprints = compileBlueprints(sampleParams)
	// fmt.Printf("%#v", route.blueprints)

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := &Context[Params]{}
		ctxV := reflect.ValueOf(ctx)

		for _, b := range route.blueprints {
			if b.GetFrom == "query" {
				rawQuery := r.URL.Query().Get(b.ParamName)
				// try to "cast" from rawQuery to the correct type (hard)
				// perhaps try to make a struct and json unmarshal to that?
			} else if b.GetFrom == "path" {
				panic("not implemented")
			} else if b.GetFrom == "body" {
				panic("not implemented")
			} else {
				fmt.Println("Unknown GetFrom option", b.GetFrom)
			}
		}

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
