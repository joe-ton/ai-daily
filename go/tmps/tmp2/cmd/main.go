package main

import (
	"context"
	"fmt"
	"os"
)

func main() {
	ctx := context.Background()

	if err := run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	fmt.Println("Hello, World!")

	<-ctx.Done()
	return nil
}
