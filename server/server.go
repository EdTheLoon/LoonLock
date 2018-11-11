package server

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// The Server struct
// keyDir			The filesystem location of keys
// assets			The filesystem location of static assets
// router			The HTTP router to be used
// fs				The FileServer for static files
type Server struct {
	keyDir string
	assets string
	router *mux.Router
	fs     http.Handler
	log    *os.File
}

// NewServer creates a new server
func NewServer(keydir string, assets string, log string) Server {
	// Create/open the log file
	f, err := os.OpenFile(log, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0640)
	if err != nil {
		panic(err)
	}

	// Create the server struct
	s := Server{
		keydir,
		assets,
		mux.NewRouter(),
		http.FileServer(http.Dir(assets)),
		f,
	}

	// Create the keys directory if it does not exist
	_, err = os.Stat(s.keyDir)
	if err != nil {
		os.MkdirAll(s.keyDir, 0700)
		s.Log("Created keys directory: " + s.keyDir)
	}

	// Initiate the routes
	s.routes()
	return s
}

// GetRouter returns the mux Router in use
func (s *Server) GetRouter() *mux.Router {
	return s.router
}
