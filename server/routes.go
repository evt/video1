package server

// routes lists routes for our HTTP server
func (s *Server) routes() {
	// index page
	s.router.HandleFunc("/blablabla", s.handleIndex())
}
