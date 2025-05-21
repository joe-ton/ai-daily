package db

import (
	"database/sql"
	"log/slog"
	"os"

	"github.com/go-sql-driver/mysql"
)

type MySQLStorage struct {
	cfg mysql.Config
	log *slog.Logger
}

func NewMySQLStorage(cfg mysql.Config, log *slog.Logger) *MySQLStorage {
	return &MySQLStorage{
		cfg: cfg,
		log: log,
	}
}

func (s *MySQLStorage) Build() (*sql.DB, error) {
	db, err := sql.Open("mysql", s.cfg.FormatDSN())
	if err != nil {
		s.log.Error("Storage setup failed", "storage", err)
		os.Exit(1) // fatal
	}
	return db, nil
}
