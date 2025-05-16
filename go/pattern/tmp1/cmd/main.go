package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

// bootstrap
func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	ctx, done := signal.NotifyContext(
		context.Background(), os.Interrupt, syscall.SIGTERM)
	defer done()

	if err := run(ctx, logger); err != nil {
		logger.Error("Application run failed", "run", err)
		os.Exit(1)
	}
}

// application
func run(ctx context.Context, logger *slog.Logger) error {
	<-ctx.Done()
	return nil
}
