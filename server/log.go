package server

import (
	"fmt"
	"time"
)

// Log a message to console and log file
func (s *server) Log(msg string) {
	newMsg := "[" + time.Now().Format(time.UnixDate) + "] " + msg
	// Print to console
	fmt.Println(newMsg)
	// Append to log file
	s.log.WriteString(newMsg + "\n")
}

func (s *server) CloseLog() {
	s.log.Close()
}
