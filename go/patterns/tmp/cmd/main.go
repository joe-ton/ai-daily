package main

type MovingAverage struct {
	size  int   // windows max size
	sum   int   // current total of integers
	queue []int // FIFO
}

func Constructor(size int) MovingAverage {
	return MovingAverage{
		size:  size,
		sum:   0,
		queue: make([]int, 0, size),
	}
}

func (m *MovingAverage) Next(val int) float64 {
	m.queue = append()
}
