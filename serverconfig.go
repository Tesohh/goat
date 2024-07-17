package goat

type OpenAPIConfig struct {
	OpenAPI    string
	Info       ServerInfo
	Paths      int // TODO:
	Components int // TODO:
}

type ServerInfo struct {
	Title          string  `json:"title"`
	Summary        string  `json:"summary,omitempty"`
	Description    string  `json:"description,omitempty"`
	TermsOfService string  `json:"terms_of_service,omitempty"`
	Contact        Contact `json:"contact,omitempty"`
	License        License `json:"license,omitempty"`
	Version        string  `json:"version"`
}

type Contact struct {
	Name  string `json:"name,omitempty"`
	URL   string `json:"url,omitempty"`
	Email string `json:"email,omitempty"`
}

type License struct {
	Name       string `json:"name"`
	Identifier string `json:"identifier,omitempty"`
	URL        string `json:"url,omitempty"`
}
