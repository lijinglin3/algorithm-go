package leetcode

import . "github.com/lijinglin3/algorithm-go/leetcode"

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Val == val {
		return head.Next
	}
	head.Next = deleteNode(head.Next, val)
	return head
}
