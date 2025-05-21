package api

import (
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joe-ton/service/user"
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

// Setup server for router and handler
func (s *Server) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userHandler := user.NewHandler()
	userHandler.RegisterRoute(subrouter)

	return http.ListenAndServe(s.addr, nil)
}
