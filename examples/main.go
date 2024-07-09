package main

import (
	"github.com/Tesohh/goat"
)

type Z struct {
	x1 int
	x2 int
}

type TestGoatType struct {
	X int    `goat:"x"`
	Y string `goat:"y,query"`
	// Z
}

func main() {
	route := goat.Route[TestGoatType, string]{
		Path:               "/tubre",
		Method:             "GET",
		Description:        "Tubreeee",
		ParamsDescriptions: map[string]string{},
		Handler: func(c *goat.Context[TestGoatType]) (int, *string, error) {
			var s = "wfw"
			return 200, &s, nil
		},
	}
	server := goat.NewServer()
	server.AddController(route)
	server.Listen(":8080")
}
