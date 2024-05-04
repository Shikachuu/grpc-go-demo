package web

import "net/http"

func (s *server) addRoutes() {
	s.mux.Handle("/", http.NotFoundHandler())
}
