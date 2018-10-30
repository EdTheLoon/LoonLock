package main

import (
	"fmt"
	"loonlock/server"
	"net/http"
)

func main() {
	// Create a new server
	webserver := server.NewServer("./keys/", "./assets")

	// Start the web server and listen on port 8080
	fmt.Println("Starting server on http://localhost:8080")
	if err := http.ListenAndServe(":8080", webserver.GetRouter()); err != nil {
		panic(err)
	}
}
