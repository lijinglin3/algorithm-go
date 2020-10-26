package leetcode

import . "github.com/lijinglin3/algorithm-go/leetcode"

func partition(head *ListNode, x int) *ListNode {
	l1, l2 := &ListNode{}, &ListNode{}
	t1, t2 := l1, l2
	for head != nil {
		tmp := head
		head = head.Next
		tmp.Next = nil

		if tmp.Val < x {
			t1.Next = tmp
			t1 = t1.Next
		} else {
			t2.Next = tmp
			t2 = t2.Next
		}
	}
	t1.Next = l2.Next
	return l1.Next
}
