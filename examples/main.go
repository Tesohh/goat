package main

import (
	"fmt"

	"github.com/Tesohh/goat"
)

type HelloBody struct {
	Title  string `json:"title"`
	Over18 bool   `json:"over18"`
}

type HelloParams struct {
	Name string `goat:"name,path"`
	HelloBody
}

var hellohtml = goat.Route[HelloParams, string]{
	Path:               "/html/{name}",
	Method:             "POST",
	Description:        "Hello.html",
	ParamsDescriptions: map[string]string{},
	Handler: func(c *goat.Context[HelloParams]) (int, *string, error) {
		var s = fmt.Sprintf("<h1>hello %s %s</h1>", c.Params.Title, c.Params.Name)
		return 200, &s, nil
	},
	OverrideEncoder: goat.HTMLEncoder,
}

var hellojson = goat.Route[HelloParams, string]{
	Path:               "/json/{name}",
	Method:             "POST",
	Description:        "Hello.json",
	ParamsDescriptions: map[string]string{},
	Handler: func(c *goat.Context[HelloParams]) (int, *string, error) {
		var s = fmt.Sprintf("hello %s %s", c.Params.Title, c.Params.Name)
		return 200, &s, nil
	},
	OverrideEncoder: goat.JSONEncoder,
}

var info = goat.ServerInfo{
	Title:          "Hello Server",
	Summary:        "Greeter Edge GenAI AR Blockchain Crypto NFT Microservice as a Service",
	Description:    "Innovative Microservice that uses the power of GenAI to greet users",
	TermsOfService: "https://example.com",
	Contact: goat.Contact{
		Name:  "Elon Must",
		URL:   "https://example.com",
		Email: "elonmust@example.com",
	},
	License: goat.License{
		Name:       "Unlicense",
		Identifier: "Unlicense",
		URL:        "https://example.com",
	},
	Version: "1.0.0",
}

func main() {
	server := goat.NewServer(info)
	server.AddController(hellohtml)
	server.AddController(hellojson)
	server.Listen(":8080")
}
