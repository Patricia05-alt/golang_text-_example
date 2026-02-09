package main

import (
	"fmt"
	"net/http" // The standard library for HTTP client/server
)

// handler receives the request and writes the response
func handler(w http.ResponseWriter, r *http.Request) {
	// 'w' is where we write the response data
	// 'r' contains details about the incoming request
	fmt.Fprintf(w, "Hello, World! This is a Go HTTP Server.")
}

func main() {
	// 1. Register the 'handler' function to the root path "/"
	http.HandleFunc("/", handler)

	// 2. Define the port
	port := ":8080"
	fmt.Println("Hello Patricia Server is running on http://localhost" + port)

	// 3. Start the server
	// ListenAndServe blocks the program and keeps it running
	err := http.ListenAndServe(port, nil)

	// Error handling: if the server fails to start (e.g., port in use)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
