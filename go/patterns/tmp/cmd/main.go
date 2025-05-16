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
		os.Exit(1)
	}
}

// application
// returns only error because of long-running
func run(ctx context.Context, logger *slog.Logger) error {
	// TODO: guard clauses

	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		logger.Error("Application failed", "run", err)
		os.Exit(1)
	}

	<-ctx.Done()
	return nil
}
