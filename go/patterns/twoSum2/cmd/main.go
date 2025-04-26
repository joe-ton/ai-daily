package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/joe-ton/leetcode"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	ctx, stop := signal.NotifyContext(
		context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := run(ctx, logger); err != nil {
		logger.Error("run application stopped", "run", err)
	}
}

func run(ctx context.Context, logger *slog.Logger) (err error) {
	nums := []int{1, 2, 3, 4}
	target := 7

	twoSum := leetcode.TwoSum{Nums: nums, Target: target}
	resp, err := twoSum.Find()
	if err != nil {
		logger.Error("run application", "run", err)
	}
	logger.Info("TwoSum output", "TwoSum", resp)

	<-ctx.Done()
	return nil
}
