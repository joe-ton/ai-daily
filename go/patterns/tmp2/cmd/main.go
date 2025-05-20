package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"

	"github.com/go-sql-driver/mysql"
	"github.com/joe-ton/db"
	"github.com/joe-ton/internal/config"
)

func main() {
	ctx := context.Background()

	if err := run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	storage := db.NewMySQLStorage(mysql.Config{
		User:      config.Envs.DBUser,
		Passwd:    config.Envs.DBPassword,
		Addr:      config.Envs.DBAddr,
		DBName:    config.Envs.DBName,
		Net:       "tcp",
		ParseTime: true,
	}, logger)
	db, err := storage.Build()
	if err != nil {
		logger.Error("Storage setup failed", "storage", err)
		os.Exit(1) // fatal
	}

	initStorage(db)

	<-ctx.Done()
	logger.Info("App gracefully shutdown", "app", "shutdown")
	return nil
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("storage success")
}
