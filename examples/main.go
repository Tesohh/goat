package main

import (
	"fmt"

	"github.com/Tesohh/goat"
)

type HelloParams struct {
	Name string
}

var hellohtml = goat.Route[HelloParams, string]{
	Path:               "/html",
	Method:             "GET",
	Description:        "Hello.html",
	ParamsDescriptions: map[string]string{},
	Handler: func(c *goat.Context[HelloParams]) (int, *string, error) {
		var s = fmt.Sprintf("<h1>hello %s</h1>", c.Params.Name)
		return 200, &s, nil
	},
	OverrideEncoder: goat.HTMLEncoder,
}

var hellojson = goat.Route[HelloParams, string]{
	Path:               "/json",
	Method:             "GET",
	Description:        "Hello.json",
	ParamsDescriptions: map[string]string{},
	Handler: func(c *goat.Context[HelloParams]) (int, *string, error) {
		var s = fmt.Sprintf("hello %s", c.Params.Name)
		return 200, &s, nil
	},
	OverrideEncoder: goat.JSONEncoder,
}

func main() {
	server := goat.NewServer()
	server.AddController(hellohtml)
	server.AddController(hellojson)
	server.Listen(":8080")
}
