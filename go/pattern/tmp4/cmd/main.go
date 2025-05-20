package main

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/signal"
)

func main() {
	ctx := context.Background()

	if err := run(
		ctx,
		os.Getenv,
		os.Args,
		os.Stdin,
		os.Stdout,
		os.Stderr,
	); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(
	ctx context.Context,
	getenv func(string) string,
	args []string,
	stdin io.Reader,
	stdout, stderr io.Writer,
) error {
	dsn := getenv
	logger := slog.New(slog.NewJSONHandler(stdout, nil))
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	<-ctx.Done()
	logger.Info("Application shutdown gracefully", "run", "shutdown")
	return nil
}
