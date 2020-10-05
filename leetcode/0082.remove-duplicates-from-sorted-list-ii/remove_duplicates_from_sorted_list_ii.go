package leetcode

import . "github.com/lijinglin3/algorithm-go/leetcode"

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Next: head}

	current := dummy
	var delvar int
	for current.Next != nil && current.Next.Next != nil {
		if current.Next.Val == current.Next.Next.Val {
			delvar = current.Next.Val
			for current.Next != nil && current.Next.Val == delvar {
				current.Next = current.Next.Next
			}
		} else {
			current = current.Next
		}

	}
	return dummy.Next
}
