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
		logger.Error("Application failed run", "run", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, logger *slog.Logger) error {
	// server listens to the port
	server := api.NewAPIServer(":8080", nil)
	resp, err := server.Run()
	if err != nil {
		logger.Error("Server port failed", "server", err)
		os.Exit(1) // fatal
	}

	<-ctx.Done()
	return nil
}
