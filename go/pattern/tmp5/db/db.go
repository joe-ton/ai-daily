package db

import (
	"database/sql"
	"log/slog"
	"os"

	"github.com/go-sql-driver/mysql"
)

type SQLStorage struct {
	cfg mysql.Config
	log *slog.Logger
}

func NewMySQLSTorage(cfg mysql.Config, log *slog.Logger) *SQLStorage {
	return &SQLStorage{
		cfg: cfg,
		log: log,
	}
}

func (s *SQLStorage) Build() (*sql.DB, error) {
	db, err := sql.Open("mysql", s.cfg.FormatDSN())
	if err != nil {
		s.log.Error("SQL Storage build failed", "sql", err)
		os.Exit(1)
	}
	return db, nil
}
