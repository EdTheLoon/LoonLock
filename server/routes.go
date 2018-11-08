package server

import "net/http"

func (s *server) routes() {

	s.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", s.fs))

	// Web Interface handling
	s.router.HandleFunc("/web/{page}", s.viewHandler)                       // Handles showing all front end HTML pages
	s.router.HandleFunc("/keys/{id}", s.adminOnly(s.getKey)).Methods("GET") // modifykey.html

	// Handles logging in or out
	s.router.HandleFunc("/login", s.loginHandler).Methods("POST")

	// Admin page handler
	s.router.HandleFunc("/admin", s.adminOnly(s.admin))

	// Key Handlers
	s.router.HandleFunc("/keys/add/{id}", s.adminOnly(s.addKey)).Methods("POST") // When key is added in addkey.html
	s.router.HandleFunc("/keys/delete/{id}", s.adminOnly(s.deleteKey))           // when 'delete' is clicked on keys.html or modifykeys.html

	// Door lock handlers
	s.router.HandleFunc("/door/lock", s.adminOnly(s.lockDoor))
	s.router.HandleFunc("/door/unlock", s.adminOnly(s.unlockDoor))
	s.router.HandleFunc("/door/unlock/{time}", s.adminOnly(s.unlockDoorTemp))
}
