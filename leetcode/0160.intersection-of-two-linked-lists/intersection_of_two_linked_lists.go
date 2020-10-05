package leetcode

import . "github.com/lijinglin3/algorithm-go/leetcode"

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	len1, len2 := 0, 0
	p1, p2 := headA, headB
	for p1 != nil {
		len1++
		p1 = p1.Next
	}
	for p2 != nil {
		len2++
		p2 = p2.Next
	}

	p1, p2 = headA, headB
	if len1 > len2 {
		for i := 0; i < len1-len2; i++ {
			p1 = p1.Next
		}
	} else {
		for i := 0; i < len2-len1; i++ {
			p2 = p2.Next
		}
	}

	for p1 != nil || p2 != nil {
		if p1 == p2 {
			return p1
		}

		p1, p2 = p1.Next, p2.Next
	}
	return nil
}
