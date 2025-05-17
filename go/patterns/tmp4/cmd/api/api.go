// cmd/api/api.go

package api

import (
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joe-ton/service/user"
)

// setup API Server struct
type APIServer struct {
	addr string
	db   *sql.DB
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
	subrouter := router.PathPrefix("api/v1").Subrouter()

	userHandler := user.NewHandler()
	userHandler.RegisterRoute(subrouter)

	s.log.Info("Router of", "addr", s.addr)

	// TODO: insert handler
	return http.ListenAndServe(s.addr, nil)
}
