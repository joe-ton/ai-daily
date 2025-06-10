package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/joe-ton/solution"
)

func main() {
	ctx := context.Background()

	if err := run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer cancel()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	nums := []int{1, 2, 3, 3}
	result, err := solution.Is_duplicate(nums)
	if err != nil {
		logger.Error("Application run error", "run", err)
		os.Exit(1)
	}
	logger.Info("Application run result", "run", result)

	<-ctx.Done()
	return nil
}
