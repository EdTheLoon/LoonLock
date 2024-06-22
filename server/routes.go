package server

import "net/http"

func (s *Server) routes() {
	// Serves static asset files through a fileserver
	s.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", s.fs))

	// Web Interface handling
	s.router.HandleFunc("/", s.viewHandler)
	s.router.HandleFunc("/web/", s.viewHandler)
	s.router.HandleFunc("/web/admin/{page}", s.adminOnly(s.viewHandler)) // Handles showing all front end HTML pages

	// Handles logging in or out
	s.router.HandleFunc("/login", s.loginHandler)

	// Admin page handler
	s.router.HandleFunc("/admin", s.adminOnly(s.admin))

	// Key Handlers
	s.router.HandleFunc("/keys/add", s.adminOnly(s.addKey)).Methods("POST")            // When key is added in addkey.html
	s.router.HandleFunc("/keys/delete/{id}", s.adminOnly(s.deleteKey))                 // when 'delete' is clicked on keys.html or modifykeys.html
	s.router.HandleFunc("/keys/{id}", s.adminOnly(s.getKey)).Methods("GET")            // Show the modify key page
	s.router.HandleFunc("/keys/update/{id}", s.adminOnly(s.updateKey)).Methods("POST") // Update a key
	s.router.HandleFunc("/keys/list", s.adminOnly(s.getAllKeys))                       // List all keys

	// Door lock handlers
	s.router.HandleFunc("/door/lock", s.adminOnly(s.lockDoor))
	s.router.HandleFunc("/door/unlock", s.adminOnly(s.unlockDoor))
	s.router.HandleFunc("/door/unlock/{time}", s.adminOnly(s.unlockDoorTemp))
}
