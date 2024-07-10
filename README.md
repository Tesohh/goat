# goat
self documenting web framework for Go

# Why goat?
I was mesmerized by ASP.NET Core's ability to automatically generate Swagger specs, manage query/body/path params automatically, but I didn't want to cheat on Go by using C#, which I personally don't like writing, so here we are.

# How does it work?
Every route is a value of the `goat.Route` struct.
The struct contains everything that goat needs to serve on that route, build the params on demand and document itself using Swagger/OpenAPI.
Under the hood it uses `net/http` to serve and [`swagno`](https://github.com/go-swagno/swagno) to generate the documentation.
Most of the reflection is done at "startup-time", while the "request-time" reflection is minimal.

# Usage (draft)
```go
package main

type HelloBody struct {
	Title string `json:"title"`
}

type HelloParams struct {
	Name string // by default, will get name from the query parameters
	Surname string `goat:"surname,path"` // if you want to get it from the path, you need to specify it with a struct tag
	HelloBody // Struct embedding works great here. This will be parsed as JSON from the body of the request.
}

helloHandler := goat.Route[HelloParams, string] {
	Route: "/{surname}",
	Method: http.MethodGET,
	Description: "Greets someone",
	Params: HelloParams {},
	ParamsDescriptions: {
		"Name": "Name of the person to greet",
	},
	Handler: func (c *goat.Context[HelloParams]) (goat.Response[string], error) {
		return goat.NewResponse(200, fmt.Sprintf("Hello, %s %s!", c.Params.Title, c.Params.Name)), nil
	}
}

func main() {
	s := goat.NewServer()
	s.AddHandler(helloHandler)
	s.AddSwagger("/swagger")
	s.Serve(":8080")
}
```
