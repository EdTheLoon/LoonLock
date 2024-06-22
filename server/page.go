package server

import (
	"os"
)

func (s *Server) loadPage(p string) ([]byte, error) {
	page, err := os.ReadFile(s.assets + p)
	return page, err
}
