package api

import (
	"database/sql"
	"log/slog"
)

type Server struct {
	addr string
	db   *sql.DB
	log  *slog.Logger
}
