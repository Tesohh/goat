package goat

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func DefaultErrorHandler(w http.ResponseWriter, status int, err error) {
	// TODO: improve
	w.WriteHeader(status)
	fmt.Fprintf(w, "%s", err.Error())
}

func JSONEncoder(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

func HTMLEncoder(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "text/html")
	s, ok := v.(string)
	if !ok {
		fmt.Fprint(w, "value provided to HTMLEncoder is not a string")
		return
	}

	fmt.Fprint(w, s)
}
