package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	tp, err := initTracer()

	if err := run(ctx, logger); err != nil {
		logger.Error("Run Error", "run", err)
	}
}

func run(ctx context.Context, logger *slog.Logger) (err error) {
}
