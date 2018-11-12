package server

import (
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
)

const cookieName string = "LoonLock"

var validSessions []Session

// Session struct for keeping the user session
type Session struct {
	ID      uuid.UUID
	Valid   bool
	Expires time.Time
	Created time.Time
	IP      string
	Agent   string
}

// sessionNew creates a new session
func (s *Server) sessionNew(ipadd string, uagent string) (*Session, error) {

	newID, err := uuid.NewRandom()
	if err != nil {
		s.Log("Error creating a new session: " + err.Error())
		return nil, err
	}

	sess := Session{
		ID:      newID,
		Valid:   true,
		Expires: time.Now().AddDate(0, 0, 1),
		Created: time.Now(),
		IP:      ipadd,
		Agent:   uagent,
	}
	validSessions = append(validSessions, sess)
	// If all executed successfully, return the session and no error
	return &sess, nil
}

func (s *Server) setSessionCookie(w http.ResponseWriter, r *http.Request) error {
	// Create a new session
	sess, err := s.sessionNew(r.RemoteAddr, r.UserAgent())

	if err != nil {
		return err
	}

	cookie := &http.Cookie{
		Name:     cookieName, // const of your app name
		Value:    sess.ID.String(),
		HttpOnly: true,
		Path:     "/",
		Secure:   false, // global var set if running ssl
		Expires:  sess.Expires,
	}

	// Write the cookie to the remote user
	http.SetCookie(w, cookie)
	return nil
}

func validSession(r *http.Request) error {
	cookies := r.Cookies()
	cValue := ""
	for i := range cookies {
		if cookies[i].Name == cookieName {
			if cookies[i].Value != "" {
				cValue = cookies[i].Value
			}
		}
	}

	if cValue == "" {
		return errors.New("Cookie value was nil")
	}

	cookieSessID, err := uuid.Parse(cValue)
	if err != nil {
		return errors.New("Cookie value was invalid")
	}

	for x := range validSessions {
		if cookieSessID == validSessions[x].ID {
			if validSessions[x].Expires.Before(time.Now()) {
				return errors.New("Session expired")
			}
			return nil
		}
	}
	return errors.New("Session was not valid")
}
