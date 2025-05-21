package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"

	"github.com/go-sql-driver/mysql"
	"github.com/joe-ton/db"
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

	db.NewMySQLStorage(mysql.Config{}, logger)

	<-ctx.Done()
	logger.Info("App graceful shutdown", "app", "shutdown")
	return nil
}
