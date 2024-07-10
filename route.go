package goat

import (
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

	OverrideErrorHandler ErrorHandlerFunc
	OverrideEncoder      EncoderFunc

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

	return func(w http.ResponseWriter, r *http.Request) {
		// Reflect (if any blueprints exist)
		var params reflect.Value
		if len(route.blueprints) > 0 { // optimisation
			params = reflect.ValueOf(&sampleParams).Elem()
		}

		for _, b := range route.blueprints {
			err := b.SetField(params, s, r)
			if err != nil {
				if route.OverrideErrorHandler != nil {
					route.OverrideErrorHandler(w, 400, err)
				} else {
					s.ErrorHandler(w, 400, err)
				}
				return
			}
		}

		// Run route
		ctx := Context[Params]{Params: sampleParams, Request: r, Writer: w}
		status, v, err := route.Handler(&ctx)

		if err != nil {
			if route.OverrideErrorHandler != nil {
				route.OverrideErrorHandler(w, 400, err)
			} else {
				s.ErrorHandler(w, status, err)
			}
			return
		}

		// Respond
		w.WriteHeader(status)
		if v != nil {
			if route.OverrideEncoder != nil {
				route.OverrideEncoder(w, *v)
			} else {
				s.Encoder(w, *v)
			}
		}

	}
}

func GenerateSwagger() ([]byte, error) {
	panic("not implemented")
}
