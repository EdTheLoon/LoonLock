package server

import (
	"io/ioutil"
)

func (s *server) loadPage(p string) ([]byte, error) {
	page, err := ioutil.ReadFile("/assets/html/" + p + ".html")
	return page, err
}
