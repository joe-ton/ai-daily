package db

import (
	"log/slog"

	"github.com/go-sql-driver/mysql"
)

type MySQLStorage struct {
	cfg mysql.Config
	log *slog.Logger
}
