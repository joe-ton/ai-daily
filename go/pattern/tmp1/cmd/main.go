package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

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
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer cancel()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	storage := db.NewMySQLStorage(mysql.Config{
		User: config.Load().DBUser,
	}, logger)
	storage.Build()

	<-ctx.Done()
	logger.Info("Application shutdown gracefully", "app", "shutdown")
	return nil
}
