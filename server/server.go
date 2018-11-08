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
type server struct {
	keyDir string
	assets string
	router *mux.Router
	fs     http.Handler
	log    *os.File
}

// NewServer creates a new server
func NewServer(keydir string, assets string, log string) server {
	f, err := os.OpenFile(log, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	s := server{
		keydir,
		assets,
		mux.NewRouter(),
		http.FileServer(http.Dir(assets)),
		f,
	}
	s.routes()
	return s
}

func (s *server) GetRouter() *mux.Router {
	return s.router
}
