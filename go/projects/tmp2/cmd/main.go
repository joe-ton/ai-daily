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
	}
}

func run(ctx context.Context, logger *slog.Logger) error {
	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		logger.Error("Server ran failed", "server", err)
	}

	<-ctx.Done()
	return nil
}
