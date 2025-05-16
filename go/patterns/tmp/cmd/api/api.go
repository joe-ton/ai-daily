package api

import (
	"database/sql"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string  // port that server listens to
	db   *sql.DB // connection pool
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

// only returns error because it is long-running
func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	return http.ListenAndServer(s.addr, router)
}
