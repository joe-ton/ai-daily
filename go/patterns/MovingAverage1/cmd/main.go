package main

type MovingAverage struct {
	size  int   // max windows size
	queue []int // slice to store last N values
	sum   int   // equals the queue
}

func Constructor(size int) MovingAverage {
	return MovingAverage{
		size:  size,
		queue: make([]int, 0, size),
		sum:   0,
	}
}
