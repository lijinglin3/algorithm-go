package leetcode

import (
	"container/heap"

	. "github.com/lijinglin2019/algorithm-go/leetcode"
)

func mergeKLists(lists []*ListNode) *ListNode {
	h := minHeap{}
	heap.Init(&h)
	node := &ListNode{}
	cur := node
	for i := 0; i < len(lists); i++ {
		if lists[i] == nil {
			continue
		}
		heap.Push(&h, lists[i])
	}

	for h.Len() > 0 {
		min := heap.Pop(&h).(*ListNode)
		cur.Next = min
		cur = cur.Next
		if min.Next != nil {
			heap.Push(&h, min.Next)
		}
	}
	return node.Next
}

type minHeap []*ListNode

func (h *minHeap) Len() int           { return len(*h) }
func (h *minHeap) Less(i, j int) bool { return (*h)[i].Val < (*h)[j].Val }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(*ListNode)) }
func (h *minHeap) Pop() interface{}   { tmp := (*h)[len(*h)-1]; *h = (*h)[:len(*h)-1]; return tmp }
func (h *minHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
