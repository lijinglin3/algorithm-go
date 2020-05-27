package leetcode

// MovingAverage ...
type MovingAverage struct {
	data []int
	size int
	sum  int
}

// Constructor ...
func Constructor(size int) MovingAverage {
	return MovingAverage{data: make([]int, 0, size), size: size}
}

// Next ...
func (a *MovingAverage) Next(val int) float64 {
	a.sum += val
	a.data = append(a.data, val)
	if len(a.data) > a.size {
		a.sum -= a.data[0]
		a.data = a.data[1:]
	}
	return float64(a.sum) / float64(len(a.data))
}
