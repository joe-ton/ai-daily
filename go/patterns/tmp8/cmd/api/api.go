package api

import (
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	addr string
	db   *sql.DB
	log  *slog.Logger
}

func NewServer(addr string, db *sql.DB, log *slog.Logger) *Server {
	return &Server{
		addr: addr,
		db:   db,
		log:  log,
	}
}

func (s *Server) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userHandler := user.Handler()
	userHandler.RegisterRoute(subrouter)

	return http.ListenAndServe(s.addr, nil)
}
