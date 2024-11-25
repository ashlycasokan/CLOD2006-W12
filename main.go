package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/echo", echoHandler)
	http.ListenAndServe(":8080", nil)
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	// Insecure: No input validation
	query := r.URL.Query().Get("input")
	fmt.Fprintf(w, "Echo: %s", query)
}

// Insecure: Exposing sensitive data via environment variable
func getEnvVariable() string {
	return os.Getenv("SECRET_KEY") // No validation or encryption
}
