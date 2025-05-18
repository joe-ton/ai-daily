package api

import (
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joe-ton/service/user"
)

type APIServer struct {
	addr string  // port
	db   *sql.DB // connection pool
	log  *slog.Logger
}

func NewAPIServer(addr string, db *sql.DB, log *slog.Logger) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
		log:  log,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	// TODO: subrouter
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userHandler := user.NewHandler()
	userHandler.RegisterRoute(subrouter)

	s.log.Info("Router updated", "addr", s.addr)
	return http.ListenAndServe(s.addr, nil)
}
