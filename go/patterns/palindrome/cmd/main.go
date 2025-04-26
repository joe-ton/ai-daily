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

	if err := run(ctx, logger); err != nil {
		logger.Error("run application stopped", "run", err)
	}
}

func run(ctx context.Context, logger *slog.Logger) (err error) {
	palindrome := solution.Palindrome{Str: "racecar"}
	resp := palindrome.IsPalindrome()

	logger.Info("Palindrome Response", "palindrome", resp)

	<-ctx.Done()
	return nil
}
