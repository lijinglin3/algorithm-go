package leetcode

import . "github.com/lijinglin2019/algorithm-go/leetcode"

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	result, tmp := &ListNode{}, &ListNode{}
	cur := result
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tmp = l1
			l1 = l1.Next
		} else {
			tmp = l2
			l2 = l2.Next
		}

		tmp.Next = nil
		cur.Next = tmp
		cur = cur.Next
	}

	if l1 == nil {
		cur.Next = l2
	} else {
		cur.Next = l1
	}

	return result.Next
}
