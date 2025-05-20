package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"

	"github.com/joe-ton/cmd/api"
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

	logger.Debug("Testing....", "debug", "debug")

	server := api.NewServer(":8080", nil, logger)

	go func() {
		if err := server.Run(); err != nil {
			logger.Error("Application server setup failed", "server", err)
			os.Exit(1)
		}
	}()

	<-ctx.Done()
	logger.Info("Graceful shutdown", "app", "shutdown")
	return nil
}
