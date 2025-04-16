package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	nums := []int{1, 2, 3, 4}
	target := 7 // return []int{2, 3}

	select {
	case <-ctx.Done(): // channel: receiver
		log.Println("Shutdown signal received...")
	default:

	}

}
