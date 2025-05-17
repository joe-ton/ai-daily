package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/joe-ton/cmd/api"
)

// bootstrap
func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	ctx, done := signal.NotifyContext(
		context.Background(), os.Interrupt, syscall.SIGTERM)
	defer done()

	if err := run(ctx, logger); err != nil {
		logger.Error("Application run failed", "run", err)
	}
}

// application
func run(ctx context.Context, logger *slog.Logger) error {
	server := api.NewAPIServer(":8080", nil, logger)
	if err := server.Run(); err != nil {
		logger.Error("Server setup failed", "server", err)
	}
	<-ctx.Done()
	return nil
}
