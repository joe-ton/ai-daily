package main

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/signal"

	"github.com/go-sql-driver/mysql"
	"github.com/joe-ton/cmd/api"
	"github.com/joe-ton/db"
	"github.com/joe-ton/internal/config"
)

func main() {
	ctx := context.Background()

	if err := run(ctx, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, stdout io.Writer) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	logger := slog.New(slog.NewJSONHandler(stdout, nil))

	storage := db.NewMySQLStorage(mysql.Config{
		User:   config.Envs.DBName,
		Passwd: "abc",
		Net:    "tcp",
		Addr:   "port",
	}, logger)
	storage.Build()
	server := api.NewServer(":8080", nil, logger)
	if err := server.Run(); err != nil {
		logger.Error("Server run failed", "server", err, "run", err)
	}
	<-ctx.Done()
	logger.Info("App shutdown gracefully", "app", "shutdown")
	return nil
}
