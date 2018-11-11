package server

import (
	"fmt"
	"time"
)

// Log a message to console and log file
func (s *Server) Log(msg string) {
	newMsg := "[" + time.Now().Format(time.UnixDate) + "] " + msg
	// Print to console
	fmt.Println(newMsg)
	// Append to log file
	s.log.WriteString(newMsg + "\n")
}

// CloseLog closes the log file
func (s *Server) CloseLog() {
	s.log.Close()
}
