package leetcode

import . "github.com/lijinglin3/algorithm-go/leetcode"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1.Next == nil && l1.Val == 0 {
		return l2
	}
	if l2.Next == nil && l2.Val == 0 {
		return l1
	}

	flag := 0
	result := &ListNode{}
	cur := result
	for l1 != nil || l2 != nil {
		var val int
		switch {
		case l1 != nil && l2 != nil:
			val = l1.Val + l2.Val + flag
			l1, l2 = l1.Next, l2.Next
		case l1 == nil && l2 != nil:
			val = l2.Val + flag
			l2 = l2.Next
		case l1 != nil && l2 == nil:
			val = l1.Val + flag
			l1 = l1.Next
		}
		if val >= 10 {
			val, flag = val-10, 1
		} else {
			flag = 0
		}
		cur.Next = &ListNode{Val: val}
		cur = cur.Next
	}
	if flag == 1 {
		cur.Next = &ListNode{Val: flag}
	}
	return result.Next
}
