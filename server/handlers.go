package server

import (
	"fmt"
	"html/template"
	"loonlock/lock"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func (s *Server) viewHandler(w http.ResponseWriter, r *http.Request) {
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

func (s *Server) loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		s.Log("Received 'login' form")
		// Read the data submitted through the form
		r.ParseForm()

		// DEBUG CODE - Comment out when not needed
		// for _k, _v := range r.Form {
		// 	fmt.Printf("%s = %s\n", _k, _v)
		// }

		user := r.FormValue("username")
		passwd := r.FormValue("password")

		if user == "admin" && passwd == "P@55w0rd" {
			s.setSessionCookie(w, r)
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			http.Error(w, "Incorrect login details", http.StatusForbidden)
			return
		}
	} else {
		t, err := template.ParseFiles("./assets/html/login.html")
		if err != nil {
			http.Error(w, "An internal server error occured", http.StatusInternalServerError)
			return
		}
		t.Execute(w, nil)
	}
}

func (s *Server) admin(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) addKey(w http.ResponseWriter, r *http.Request) {
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

	// Parse expiry input
	var err error
	var expiresTime time.Time
	var expiresStr string
	if expires != "" {
		expiresTime, err = StrToTime(expires)
		expiresStr = TimeToStr(expiresTime)
	} else {
		expiresStr = "Jan, 01 2150"
	}
	if err != nil {
		http.Error(w, "Could not parse date:\n"+err.Error(), http.StatusInternalServerError)
		s.Log("Could not parse date: " + err.Error())
		return
	}

	// Parse single use
	singleUse := false
	if singleUseStr == "on" {
		singleUse = true
	}

	s.Log("Creating key...")
	// Create the key using the provided data
	key, err := createKey(name, description, expiresStr, singleUse)
	if err != nil {
		http.Error(w, "Could not create key:\n"+err.Error(), http.StatusUnauthorized)
		s.Log("Could not create key: " + err.Error())
		return
	}

	// Write the key to file
	s.Log("Writing key...")
	err = s.writeKey(&key)
	if err != nil {
		http.Error(w, "Could not write key:\n"+err.Error(), http.StatusInternalServerError)
		s.Log("Could not write key: " + err.Error())
		return
	}
	s.Log("Key written successfully!")
	_, _ = w.Write([]byte("Success!\nKey added successfully!"))

}

func (s *Server) getKey(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) getAllKeys(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) updateKey(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) deleteKey(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) unlockDoor(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) unlockDoorTemp(w http.ResponseWriter, r *http.Request) {
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

func (s *Server) lockDoor(w http.ResponseWriter, r *http.Request) {

}

func (s *Server) adminOnly(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := validSession(r)
		if err != nil {
			if err != nil {
				s.Log(err.Error())
			}
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		next.ServeHTTP(w, r)
	}
}
