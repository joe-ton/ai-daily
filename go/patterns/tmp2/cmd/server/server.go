package server

import (
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	addr string
	db   *sql.DB
	log  *slog.Logger
}

func NewServer(addr string, db *sql.DB, log *slog.Logger) *server {
	return &server{
		addr: addr,
		db:   db,
		log:  log,
	}
}

func (s *server) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userHandler := user.NewHandler()
	userHandler.RegisterRoute(subrouter)

	s.log.Info("Router handler setup", "addr", s.addr)

	return http.ListenAndServe(s.addr, nil)
}
