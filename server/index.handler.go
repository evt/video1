package server

import (
	"net/http"
)

// handleIndex handles GET / request
func (s *Server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, map[string]interface{}{
			"version": "0.1",
			"name":    "My cool web service",
		}, 200)
	}
}
