package leetcode

import . "github.com/lijinglin3/algorithm-go/leetcode"

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{Next: head}
	sorted, curr := head, head.Next
	for curr != nil {
		if sorted.Val <= curr.Val {
			sorted = sorted.Next
		} else {
			prev := dummy
			for prev.Next.Val <= curr.Val {
				prev = prev.Next
			}
			sorted.Next = curr.Next
			curr.Next = prev.Next
			prev.Next = curr
		}
		curr = sorted.Next
	}
	return dummy.Next
}
