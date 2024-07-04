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

```
