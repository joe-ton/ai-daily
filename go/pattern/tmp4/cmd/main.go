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
		fmt.Fprintf(os.Stdout, "%s\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer cancel()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	nums := []int{1, 2, 3, 4}
	target := 7

	twoSum := solution.TwoSum{Nums: nums, Target: target}
	resp, err := twoSum.Find()
	if err != nil {
		logger.Error("Run application error", "run", err)
		os.Exit(1)
	}
	logger.Info("Run application success", "output", resp)

	<-ctx.Done()
	logger.Info("Application shutdown gracefuly...", "run", "graceful")
	return nil
}
