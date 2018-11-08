package server

import (
	"fmt"
)

// Log a message to console and log file
func (s *server) Log(msg string) {
	// Print to console
	fmt.Println(msg)
	// Append to log file
	s.log.WriteString(msg + "\n")
}

func (s *server) CloseLog() {
	s.log.Close()
}
