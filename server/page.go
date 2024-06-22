package server

import (
	"os"
)

func (s *Server) loadPage(p string) ([]byte, error) {
	page, err := os.ReadFile(s.assets + "/html/" + p)
	return page, err
}
