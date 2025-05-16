package api

import (
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string  // listen to port
	db   *sql.DB // database
	log  *slog.Logger
}

func NewAPIServer(addr string, db *sql.DB, log *slog.Logger) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
		log:  log,
	}
}

func (s APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	s.log.Info("Router setup", "router", subrouter)

	return http.ListenAndServe(s.addr, nil)
}
