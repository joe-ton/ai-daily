package api

import (
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joe-ton/service/user"
)

type server struct {
	addr string  // port; server listening
	db   *sql.DB // connection pool
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
	s.log.Info("Router setup", "addr", s.addr)

	return http.ListenAndServe(s.addr, nil)
}
