package leetcode

import . "github.com/lijinglin2019/algorithm-go/leetcode"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1.Next == nil && l1.Val == 0 {
		return l2
	}
	if l2.Next == nil && l2.Val == 0 {
		return l1
	}

	s1, s2 := toSlice(l1), toSlice(l2)
	len1, len2 := len(s1), len(s2)
	result, flag := (*ListNode)(nil), 0

	for len1 != 0 || len2 != 0 {
		val := 0
		switch {
		case len1 != 0 && len2 != 0:
			val = s1[len1-1] + s2[len2-1] + flag
			len1--
			len2--
		case len1 == 0 && len2 != 0:
			val = s2[len2-1] + flag
			len2--
		case len1 != 0 && len2 == 0:
			val = s1[len1-1] + flag
			len1--
		}
		if val >= 10 {
			val, flag = val-10, 1
		} else {
			flag = 0
		}
		result = &ListNode{Val: val, Next: result}
	}
	if flag == 1 {
		result = &ListNode{Val: flag, Next: result}
	}
	return result
}

func toSlice(l *ListNode) (result []int) {
	for l != nil {
		result = append(result, l.Val)
		l = l.Next
	}
	return
}
