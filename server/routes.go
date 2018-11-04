package server

import "net/http"

func (s *server) routes() {

	s.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", s.fs))

	// Home handler - displays a login page
	s.router.HandleFunc("/", s.homeHandler)

	// Handles logging in or out
	s.router.HandleFunc("login", s.loginHandler).Methods("POST")

	// Admin page handler
	s.router.HandleFunc("/admin", s.adminOnly(s.admin))

	// Key Handlers
	s.router.HandleFunc("/keys/add/{id}", s.adminOnly(s.addKey))
	s.router.HandleFunc("/keys/delete/{id}", s.adminOnly(s.deleteKey))
	s.router.HandleFunc("/keys/{id}", s.adminOnly(s.getKey)).Methods("GET")

	// Door lock handlers
	s.router.HandleFunc("/door/lock", s.adminOnly(s.lockDoor))
	s.router.HandleFunc("/door/unlock", s.adminOnly(s.unlockDoor))
	s.router.HandleFunc("/door/unlock/{time}", s.adminOnly(s.unlockDoorTemp))
}
