package server

import "net/http"

func (s *server) routes() {

	s.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", s.fs))

	// Key Handlers
	s.router.HandleFunc("/admin", s.adminOnly(s.admin))
	s.router.HandleFunc("/keys/add", s.adminOnly(s.addKey))
	s.router.HandleFunc("/keys/{id}", s.adminOnly(s.getKey)).Methods("GET")
	s.router.HandleFunc("/keys/{id}", s.adminOnly(s.updateKey)).Methods("PUT")
	s.router.HandleFunc("/keys/{id}", s.adminOnly(s.deleteKey)).Methods("DELETE")
}
