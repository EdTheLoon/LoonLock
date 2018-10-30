package server

import "net/http"

func (s *server) admin(w http.ResponseWriter, r *http.Request) {

}

func (s *server) addKey(w http.ResponseWriter, r *http.Request) {

}

func (s *server) getKey(w http.ResponseWriter, r *http.Request) {

}

func (s *server) updateKey(w http.ResponseWriter, r *http.Request) {

}

func (s *server) deleteKey(w http.ResponseWriter, r *http.Request) {

}

func (s *server) adminOnly(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// if !currentUser(r).IsAdmin {
		if 0 != 0 {
			http.NotFound(w, r)
			return
		}
		h(w, r)
	}
}
