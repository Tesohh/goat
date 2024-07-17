package goat

type SwaggerPath map[string]SwaggerOperation

// TODO: finish operation https://swagger.io/specification/
type SwaggerOperation struct {
	Tags        []string
	Summary     string
	Description string
}
