package leetcode

import . "github.com/lijinglin3/algorithm-go/leetcode"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	tmp := result
	flag := 0
	for l1 != nil || l2 != nil || flag != 0 {
		val := flag
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}
		flag, val = val/10, val%10
		tmp.Next = &ListNode{Val: val}
		tmp = tmp.Next
	}

	return result.Next
}
