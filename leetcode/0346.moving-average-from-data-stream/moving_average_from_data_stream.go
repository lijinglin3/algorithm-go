package leetcode

// MovingAverage ...
type MovingAverage struct {
	head, size, sum, count int
	queue                  []int
}

// Constructor ...
func Constructor(size int) MovingAverage {
	return MovingAverage{queue: make([]int, size), size: size}
}

// Next ...
func (a *MovingAverage) Next(val int) float64 {
	a.sum += val
	if a.count == a.size {
		a.sum -= a.queue[a.head]
	} else {
		a.count++
	}
	a.queue[a.head] = val
	a.head = (a.head + 1) % a.size
	return float64(a.sum) / float64(a.count)
}
