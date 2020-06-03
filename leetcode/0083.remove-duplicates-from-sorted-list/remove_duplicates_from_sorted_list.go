package leetcode

import . "github.com/lijinglin2019/algorithm-go/leetcode"

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	n := head
	for n != nil && n.Next != nil {
		if n.Val == n.Next.Val {
			n.Next = n.Next.Next
		} else {
			n = n.Next
		}
	}
	return head
}
