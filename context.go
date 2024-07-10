package goat

import "net/http"

type Context[T any] struct {
	Params  T
	Request *http.Request
	Writer  http.ResponseWriter
}
