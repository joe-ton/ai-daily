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

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := run(ctx, logger); err != nil {
		logger.Error("new", "logger", nil)
	}
}

func run(ctx context.Context, logger *slog.Logger) error {
	nums := []int{1, 2, 3, 4}
	target := 7
	twoSum := solution.TwoSum{Nums: nums, Target: target}
	resp, err := twoSum.Find()

	if err != nil {
		logger.Error("Error:", err)
		return nil
	}
	logger.Info("Response:", resp)

	<-ctx.Done()
	return nil
}
