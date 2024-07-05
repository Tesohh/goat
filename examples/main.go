package main

import "github.com/Tesohh/goat"

type Z struct {
	x1 int
	x2 int
}

type TestGoatType struct {
	X int    `goat:"x"`
	Y string `goat:"y,query"`
	Z
}

func main() {
	route := goat.Route[TestGoatType, string]{}
	route.MakeHandlerFunc(&goat.Server{})
}
