package main

import (
	"fmt"
	"log"
	"net/http" // The standard library for HTTP client/server
	"os"
	"strings"
)

// handler receives the request and writes the response
func handler(w http.ResponseWriter, r *http.Request) {
	// 'w' is where we write the response data
	// 'r' contains details about the incoming request
	fmt.Fprintf(w, "Hello, World...this Is Patricia first go project! This is a Go HTTP Server.")
}

func main() {
	// 1. Register the 'handler' function to the root path "/"
	http.HandleFunc("/", handler)

	// 2. Determine the port from environment (fallback to :8080)
	portEnv := os.Getenv("PORT")
	var port string
	if portEnv == "" {
		port = ":8080"
	} else if strings.HasPrefix(portEnv, ":") {
		port = portEnv
	} else {
		port = ":" + portEnv
	}

	fmt.Println("Hello World Server is running on http://localhost" + port)

	// 3. Start the server (blocks)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

