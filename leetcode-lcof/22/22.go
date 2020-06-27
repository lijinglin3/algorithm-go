package leetcode

import . "github.com/lijinglin2019/algorithm-go/leetcode"

func getKthFromEnd(head *ListNode, k int) *ListNode {
	p1, p2 := head, head
	for i := 0; i < k; i++ {
		p2 = p2.Next
		if p2 == nil {
			return head
		}
	}
	for p2 != nil {
		p1, p2 = p1.Next, p2.Next
	}
	return p1
}
