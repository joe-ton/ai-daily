package main

type MovingAverage struct {
	size  int // maximum window size
	queue []int
	sum   int
}
