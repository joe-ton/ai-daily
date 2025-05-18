package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/joe-ton/cmd/api"
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
	// db, err := db.NewMySQLStorage(mysql.Config{
	// 	User:                 "root",
	// 	Passwd:               "abc",
	// 	Addr:                 "127.0.1:3306",
	// 	DBName:               "econ",
	// 	Net:                  "tcp",
	// 	AllowNativePasswords: true,
	// 	ParseTime:            true,
	// })
	server := api.NewAPIServer(":8080", nil, logger)
	go func() {
		if err := server.Run(); err != nil {
			logger.Error("Application failed...", "run", err)
			os.Exit(1)
		}
	}()

	<-ctx.Done()
	return nil
}
