package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/joe-ton/ai-daily/solution"
	"go.uber.org/zap"
)

func main() {
	logger := zap.Must(zap.NewProduction())
	defer logger.Sync()
	sugar := logger.Sugar()

	ctx, stop := signal.NotifyContext(
		context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := run(ctx, sugar); err != nil {
		logger.Fatal("existing", zap.Error(err))
	}

	logger.Info("graceful shutdown")
}

func run(ctx context.Context, log *zap.SugaredLogger) error {
	select {
	case <-ctx.Done():
		return ctx.Err() // honour cancellation
	default:
	}
	nums := []int{1, 2, 3, 4}
	target := 7
	// TODO: get rid of postional-based initialization
	twoSum := solution.TwoSum{Nums: nums, Target: target}
	resp, err := twoSum.Find()
	if err != nil {
		return err
	}
	log.Infow("TwoSum solved", "response", resp)
	return nil
}
