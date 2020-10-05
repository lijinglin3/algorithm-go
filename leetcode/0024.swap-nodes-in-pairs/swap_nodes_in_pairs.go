package leetcode

import . "github.com/lijinglin3/algorithm-go/leetcode"

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur, pre := head, head.Next
	cur.Next = swapPairs(pre.Next)
	pre.Next = cur
	return pre
}
