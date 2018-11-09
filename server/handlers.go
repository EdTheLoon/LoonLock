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
	s.Log("Received 'addkey' form")
	// Read the data submitted through the form
	r.ParseForm()

	// DEBUG CODE - Comment out when not needed
	// for _k, _v := range r.Form {
	// 	fmt.Printf("%s = %s\n", _k, _v)
	// }

	name := r.FormValue("name")
	description := r.FormValue("description")
	expires := r.FormValue("expires")
	singleUseStr := r.FormValue("singleUse")
	singleUse := false
	if singleUseStr == "on" {
		singleUse = true
	}

	s.Log("Creating key...")
	// Create the key using the provided data
	key, err := createKey(name, description, expires, singleUse)
	if err != nil {
		http.Error(w, "Could not create key:\n"+err.Error(), http.StatusSeeOther)
		s.Log("Could not create key: " + err.Error())
		return
	}

	// Write the key to file
	s.Log("Writing key...")
	err = s.writeKey(&key)
	if err != nil {
		http.Error(w, "Could not write key:\n"+err.Error(), http.StatusSeeOther)
		s.Log("Could not write key: " + err.Error())
		return
	}
	s.Log("Key written successfully!")
	_, _ = w.Write([]byte("Success!\nKey added successfully!"))

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
		// TO DO: IMPLEMENT PROPER AUTHORISATION CHECKING
		if 0 != 0 { // CHANGE THIS
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	}
}
