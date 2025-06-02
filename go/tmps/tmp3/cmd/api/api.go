package api

import (
	"database/sql"
	"log/slog"
)

type Server struct {
	addr string // port for server to listen to
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

// handler + router
func (s *Server) Run() error {

}
