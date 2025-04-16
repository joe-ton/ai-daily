package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joe-ton/solution"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	nums := []int{1, 2, 3, 4}
	target := 7

	select {
	case <-ctx.Done(): // channel: receive-only type
		log.Println("Received shutdown signal. Gracefully shutdown...")
	default:
		twosum := solution.TwoSum{Nums: nums, Target: target}
		resp, err := twosum.Find()
		if err != nil {
			log.Println("Error:", err)
		} else {
			log.Println("Response:", resp)
		}
	}
}
