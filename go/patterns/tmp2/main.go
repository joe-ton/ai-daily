package main

import (
	"context"
	"fmt"
	"os"
)

func main() {
	ctx := context.Background()

	if err := run(
		ctx,
		os.Getenv,
		os.Args,
		os.Stdin,
		os.Stdout,
		os.Stderr,
	); err != nil {
		fmt.Fprintf(os.Stderr, err)
		os.Exit(1)
	}
}

func run(
    ctx context.Context,
    getenv func(string) string, 

    os.Args,
    os.Stdin,
    os.Stdout,
    os.Stderr,
)
