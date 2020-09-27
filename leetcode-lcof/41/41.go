package leetcode

import "container/heap"

// MedianFinder ...
type MedianFinder struct {
	min, max *intHeap
}

// Constructor initialize your data structure here.
func Constructor() MedianFinder {
	return MedianFinder{min: &intHeap{}, max: &intHeap{}}
}

// AddNum ...
func (f *MedianFinder) AddNum(num int) {
	if f.min.Len() > 0 && num > (*f.min)[0] {
		heap.Push(f.min, num)
	} else {
		heap.Push(f.max, -num)
	}

	switch f.min.Len() - f.max.Len() {
	case 2:
		heap.Push(f.max, -heap.Pop(f.min).(int))
	case -2:
		heap.Push(f.min, -heap.Pop(f.max).(int))
	}
}

// FindMedian ...
func (f *MedianFinder) FindMedian() float64 {
	switch {
	case f.min.Len() > f.max.Len():
		return float64((*f.min)[0])
	case f.min.Len() < f.max.Len():
		return -float64((*f.max)[0])
	default:
		return float64((*f.min)[0]-(*f.max)[0]) / 2
	}
}

type intHeap []int

func (h *intHeap) Len() int             { return len(*h) }
func (h *intHeap) Less(i, j int) bool   { return (*h)[i] < (*h)[j] }
func (h *intHeap) Swap(i, j int)        { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *intHeap) Push(val interface{}) { *h = append(*h, val.(int)) }
func (h *intHeap) Pop() interface{}     { tmp := (*h)[len(*h)-1]; *h = (*h)[:len(*h)-1]; return tmp }
