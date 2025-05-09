package main

type MovingAverage struct {
	size  int   // max window size, fixed
	queue []int // slice of last N integers
	sum   int   // total value of current N integers
}

func Constructor(size int) MovingAverage {
	return MovingAverage{
		size:  size,
		queue: make([]int, 0, size),
		sum:   size,
	}
}

func (m *MovingAverage) Next(val int) float64 {
	m.queue = append()
}

327419653001
