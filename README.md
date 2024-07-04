# goat
self documenting web framework for Go. No magic, just reflection and generics.

# Why goat?
I was mesmerized by ASP.NET Core's ability to automatically generate Swagger specs, manage query/body/path params automatically, but I didn't want to cheat on Go by using C#, which I personally don't like writing, so here we are.

# How does it work?
Every route is a value of the `goat.Route` struct.
The struct contains everything that goat needs to serve on that route, build the params on demand and document itself using Swagger/OpenAPI.
Under the hood it uses `net/http` to serve and [`swagno`](https://github.com/go-swagno/swagno) to generate the documentation.

# Usage (draft)
```go
package main

type HelloParams struct {
	Name string `goat:"name,query"`
}

helloHandler := goat.Route {
	Route: "/",
	Method: http.MethodGET,
	Description: "Greets someone",
	Params: HelloParams {},
	ParamsDescriptions: {
		"Name": "Name of the person to greet"
	},
	Handler: func (c *goat.Context[HelloParams]) (goat.Response[string], error) {
		return goat.NewResponse(200, fmt.Sprintf("Hello, %s!", c.Params.Name)), nil
	}
}

func main() {
	s := goat.NewServer()
	s.AddHandler(helloHandler)
	s.Serve(":8080")	
}
```
