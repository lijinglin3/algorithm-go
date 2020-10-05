package leetcode

import . "github.com/lijinglin3/algorithm-go/leetcode"

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p1, p2 := head, head
	for i := 0; i < n; i++ {
		p2 = p2.Next
	}

	// p2 为 nil 时表示删除的是第一个节点
	if p2 == nil {
		return head.Next
	}

	for {
		p2 = p2.Next
		if p2 == nil {
			break
		}
		p1 = p1.Next
	}

	p1.Next = p1.Next.Next

	return head
}
