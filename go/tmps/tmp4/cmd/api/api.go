package api

import (
	"database/sql"
	"log/slog"
	"net/http"
)

type Server struct {
	addr string // port
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

func (s *Server) Run() error {
	router := http.NewServeMux()

	router.HandleFunc("/login", s.handleLogin)
	router.HandleFunc("/health", s.handleHealth)

	srv := &http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	return nil
}

func (s *Server) handleLogin(w http.ResponseWriter, r *http.Request) {
}
func (s *Server) handleHealth(w http.ResponseWriter, r *http.Request) {
}
