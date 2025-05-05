package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/joe-ton/solution"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	ctx, stop := signal.NotifyContext(
		context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := run(logger, ctx); err != nil {
		logger.Error("Run application failed", "run", err)
	}
}

func run(logger *slog.Logger, ctx context.Context) (err error) {
	nums := []int{1, 2, 3, 4}
	target := 7
	twoSum := solution.TwoSum{Target: target, Nums: nums}
	resp, err := twoSum.Find()

	if err != nil {
		logger.Error("Find error", "find", err)
	}
	logger.Info("Find response", "find", resp)

	<-ctx.Done()
	return nil
}
