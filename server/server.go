package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

// The Server struct
// keyDir			The filesystem location of keys
// assets			The filesystem location of static assets
// router			The HTTP router to be used
// fs				The FileServer for static files
type server struct {
	keyDir string
	assets string
	router *mux.Router
	fs     http.Handler
}

// NewServer creates a new server
func NewServer(kd string, ass string) server {
	s := server{
		kd,
		ass,
		mux.NewRouter(),
		http.FileServer(http.Dir("./assets")),
	}

	return s
}

func (s *server) GetRouter() *mux.Router {
	return s.router
}
