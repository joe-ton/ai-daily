package api

import (
	"database/sql"

	"github.com/gorilla/mux"
)

type APIServer struct {
	address string  // port
	db      *sql.DB // pool connection
}

func NewAPIServer(address string, db *sql.DB) *APIServer {
	return &APIServer{
		address: address,
		db:      db,
	}
}

func (s APIServer) Run() error {
	router := mux.NewRouter()
}
