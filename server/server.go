package server

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/evt/video1/config"
	"github.com/evt/video1/db"
	"github.com/gorilla/mux"
)

// Server is a server with all the batteries included :)
type Server struct {
	context   context.Context
	config    *config.Config
	router    *mux.Router
	db        *db.PgDB
}

// Init returns new server instance
func Init(ctx context.Context, config *config.Config, db *db.PgDB) *Server {
	router := mux.NewRouter()
	s := &Server{
		context:   ctx,
		config:    config,
		router:    router,
		db:        db,
	}
	s.routes()
	return s
}

// respond converts data to JSON and sends it to client
func (s *Server) respond(w http.ResponseWriter, r *http.Request, data interface{}, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			// TODO
		}
	}
}

// error sends { "error": ... } to client
func (s *Server) error(w http.ResponseWriter, r *http.Request, err error, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	if err != nil {
		err := json.NewEncoder(w).Encode(e(err))
		if err != nil {
			// TODO
		}
	}
}

// decode decodes incoming JSON request
func (s *Server) decode(w http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

// ServeHTTP makes our server http.Handler
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
