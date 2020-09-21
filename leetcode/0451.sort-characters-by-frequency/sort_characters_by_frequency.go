package leetcode

import (
	"bytes"
	"container/heap"
)

func frequencySort(s string) string {
	m := make(map[byte]*node)
	h := nodeHeap{}
	for _, b := range []byte(s) {
		if n, ok := m[b]; ok {
			n.count++
		} else {
			tmp := &node{value: b, count: 1}
			m[b] = tmp
			h.Push(tmp)
		}
	}

	heap.Init(&h)
	result := make([]byte, 0, len(s))
	for h.Len() != 0 {
		n := heap.Pop(&h).(*node)
		result = append(result, bytes.Repeat([]byte{n.value}, n.count)...)
	}
	return string(result)
}

type node struct {
	value byte
	count int
}

type nodeHeap []*node

func (h *nodeHeap) Len() int           { return len(*h) }
func (h *nodeHeap) Less(i, j int) bool { return (*h)[i].count > (*h)[j].count }
func (h *nodeHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *nodeHeap) Push(x interface{}) { *h = append(*h, x.(*node)) }
func (h *nodeHeap) Pop() interface{}   { x := (*h)[len(*h)-1]; *h = (*h)[:len(*h)-1]; return x }
