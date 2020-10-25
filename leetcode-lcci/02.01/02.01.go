package leetcode

import . "github.com/lijinglin3/algorithm-go/leetcode"

func removeDuplicateNodes(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	cache := make(map[int]bool, 128)
	cache[head.Val] = true
	for tmp := head; tmp.Next != nil; {
		if cache[tmp.Next.Val] {
			tmp.Next = tmp.Next.Next
		} else {
			cache[tmp.Next.Val] = true
			tmp = tmp.Next
		}
	}
	return head
}
