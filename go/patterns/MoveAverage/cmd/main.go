package main

type MovingAverage struct {
	size  int   // maximum window size
	queue []int // slice to store the last N integers
	sum   int   // total current amount based on queue
}

func Constructor(size int) MovingAverage {
	return MovingAverage{
		size:  size,
		queue: make([]int, 0, size),
		sum:   0,
	}
}

func (m *MovingAverage) Next(val int) float64 {
	m.queue = append(m.queue, val)
}
