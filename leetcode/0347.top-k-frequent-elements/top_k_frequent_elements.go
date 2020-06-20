package leetcode

import "container/heap"

func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	for _, v := range nums {
		count[v]++
	}

	h := &nodeHeap{}
	heap.Init(h)
	for v, c := range count {
		heap.Push(h, &node{
			value: v,
			count: c,
		})
	}

	if k > len(count) {
		k = len(count)
	}
	result := make([]int, 0, k)
	for i := 0; i < k; i++ {
		result = append(result, heap.Pop(h).(*node).value)
	}
	return result
}

type node struct {
	value int
	count int
}

type nodeHeap []*node

func (h nodeHeap) Len() int           { return len(h) }
func (h nodeHeap) Less(i, j int) bool { return h[i].count > h[j].count }
func (h nodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *nodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*node))
}

func (h *nodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
