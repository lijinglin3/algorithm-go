package leetcode

import . "github.com/lijinglin2019/algorithm-go/leetcode"

func reverseList(head *ListNode) *ListNode {
	var ln *ListNode
	for head != nil {
		cur := head
		head = head.Next
		cur.Next = ln
		ln = cur
	}
	return ln
}
