package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Serve static files from the "public" directory
	http.Handle("/", http.FileServer(http.Dir("public")))

	// Get the port from the environment variable or use a default value
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Printf("Server is running on port %s\n", port)
	http.ListenAndServe(":"+port, nil)
}
