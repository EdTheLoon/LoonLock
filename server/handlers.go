package server

import (
	"fmt"
	"loonlock/lock"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *server) viewHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	load := vars["page"] + ".html"
	if load == ".html" {
		load = "home.html"
	}
	p, err := s.loadPage(load)
	s.Log("Loading page: ./assets/" + load)
	if err != nil {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}
	w.Write(p)
}

func (s *server) loginHandler(w http.ResponseWriter, r *http.Request) {

}

func (s *server) admin(w http.ResponseWriter, r *http.Request) {

}

func (s *server) addKey(w http.ResponseWriter, r *http.Request) {

}

func (s *server) getKey(w http.ResponseWriter, r *http.Request) {

}

func (s *server) getAllKeys(w http.ResponseWriter, r *http.Request) {

}

func (s *server) updateKey(w http.ResponseWriter, r *http.Request) {

}

func (s *server) deleteKey(w http.ResponseWriter, r *http.Request) {

}

func (s *server) unlockDoor(w http.ResponseWriter, r *http.Request) {

}

func (s *server) unlockDoorTemp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	timeStr := vars["time"]
	time, err := strconv.Atoi(timeStr)
	if err != nil {
		// Handle an error (the 'time' is not a number)
		fmt.Println(err)
		return
	}
	fmt.Println("Unlock for '" + timeStr + "' seconds")
	lock.UnlockTemp(time)
}

func (s *server) lockDoor(w http.ResponseWriter, r *http.Request) {

}

func (s *server) adminOnly(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// if !currentUser(r).IsAdmin {
		if 0 != 0 {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	}
}
