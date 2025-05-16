package api

import "database/sql"

type APIServer struct {
	address string  // setting up port
	db      *sql.DB // connection pool
}

func NewAPIServer(address string, db *sql.DB) *APIServer {
	return &APIServer{
		address: address,
		db:      db,
	}
}
