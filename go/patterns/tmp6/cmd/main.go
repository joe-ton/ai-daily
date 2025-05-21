package main

import (
	"context"
	"fmt"
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

	cfg := config.Load()

	store := db.NewMySQLStorage(mysql.Config{
		User:   cfg.DBUser,
		Passwd: cfg.DBPass,
		Net:    "tcp",
		Addr:   cfg.DBAddr,
		DBName: cfg.DBName,
	}, logger)

	dbConn, err := store.Build()
	if err != nil {
		logger.Error("Connection failed", "connection", err)
		return err
	}
	defer dbConn.Close()

	<-ctx.Done()
	logger.Info("Application graceful shutdown", "run", "shutdown")
	return nil

}
