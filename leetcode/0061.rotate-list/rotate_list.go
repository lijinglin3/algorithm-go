package leetcode

import . "github.com/lijinglin3/algorithm-go/leetcode"

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	pre, cur := head, head.Next
	length := 1

	for cur != nil {
		pre, cur = pre.Next, cur.Next
		length++
	}

	pre.Next, cur = head, head

	for i := length - k%length; i > 0; i-- {
		pre, cur = pre.Next, cur.Next
	}
	pre.Next = nil
	return cur
}
