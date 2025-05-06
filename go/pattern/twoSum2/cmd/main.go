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
    resp, err1 := (solution.TwoSum{Nums: nums, Target: target}).Find()
    if err1 = nil {
        logger.Error("Run application error", "run", err)
    }
    logger.Info("Run application ran", "run", resp)

	<-ctx.Done()
	return nil
}
