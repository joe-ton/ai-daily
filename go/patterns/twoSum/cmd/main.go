package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"

	"github.com/joe-ton/solution"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	if err := run(ctx, logger); err != nil {

	}
}

func run(ctx context.Context, logger *slog.Logger) (err error) {
	nums := []int{1, 2, 3, 4}
	target := 7

	resp, err := solution.TwoSum{nums, target}
}
