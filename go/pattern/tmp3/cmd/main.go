package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-sql-driver/mysql"
	"github.com/joe-ton/cmd/api"
	"github.com/joe-ton/db"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	ctx, done := signal.NotifyContext(
		context.Background(), os.Interrupt, syscall.SIGTERM)
	defer done()

	if err := run(ctx, logger); err != nil {
		logger.Error("Application failed", "run", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, logger *slog.Logger) error {
	db.NewMySQLStorage(mysql.Config{
		User:                 "root",
		Passwd:               "abc",
		Addr:                 "127.0.1:3306",
		DBName:               "econ",
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}, logger)

	server := api.NewAPIServer(":8080", nil, logger)

	go func() {
		if err := server.Run(); err != nil {
			logger.Error("Application failed", "run", err)
			os.Exit(1)
		}
	}()

	<-ctx.Done()
	logger.Info(
		"Application shutdown gracefully...", "run", "server-shutdown")
	return nil
}
