package server

import (
	"io/ioutil"
)

func (s *Server) loadPage(p string) ([]byte, error) {
	page, err := ioutil.ReadFile("./assets/html/" + p)
	return page, err
}
