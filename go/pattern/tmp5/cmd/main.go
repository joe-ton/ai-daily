package main

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/signal"

	"github.com/go-sql-driver/mysql"
	"github.com/joe-ton/cmd/api"
	"github.com/joe-ton/db"
	"github.com/joe-ton/solution"
)

func main() {
	ctx := context.Background()

	if err := run(
		ctx,
		os.Stdout,
		os.Getenv,
	); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(
	ctx context.Context,
	stdout io.Writer,
	getenv func(string) string,
) error {
	logger := slog.New(slog.NewJSONHandler(stdout, nil))

	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	db.NewMySQLSTorage(mysql.Config{
		User:   "root",
		Passwd: "abc",
		Net:    "tcp",
		DBName: "mysql",
	}, logger)

	server := api.NewServer(":8080", nil, logger)
	if err := server.Run(); err != nil {
		logger.Error("New Server failed", "server", err)
		os.Exit(1)
	}

	nums := []int{1, 2, 3, 4}
	target := 7

	resp, err := solution.NewTwoSum(nums, target).Find()
	if err != nil {
		logger.Error("TwoSum find failed", "twosum", err)
		os.Exit(1)
	}
	logger.Info("TwoSum find", "run", resp)

	<-ctx.Done()
	logger.Info("App shutdown gracefully...", "app", "shutdown")
	return nil
}
