package main

import (
	"loonlock/server"
	"net/http"
)

func main() {
	// Create a new server
	s := server.NewServer("./.keys/", "./assets", ".log")
	defer s.CloseLog()

	// Start the web server and listen on port 8080
	s.Log("Server started on http://localhost:8080")
	if err := http.ListenAndServe(":8080", s.GetRouter()); err != nil {
		panic(err)
	}
}
