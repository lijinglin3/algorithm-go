package leetcode

import . "github.com/lijinglin2019/algorithm-go/leetcode"

//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	num1, num2 := 0, 0
//	for l1 != nil {
//		num1 = num1*10 + l1.Val
//		l1 = l1.Next
//	}
//	for l2 != nil {
//		num2 = num2*10 + l2.Val
//		l2 = l2.Next
//	}
//
//	sum := num1 + num2
//	result := &ListNode{Val: sum % 10}
//	sum = sum / 10
//
//	for sum != 0 {
//		result = &ListNode{Val: sum % 10, Next: result}
//		sum = sum / 10
//	}
//	return result
//}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1.Next == nil && l1.Val == 0 {
		return l2
	}
	if l2.Next == nil && l2.Val == 0 {
		return l1
	}

	var s1, s2 []int
	for l1 != nil {
		s1 = append(s1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		s2 = append(s2, l2.Val)
		l2 = l2.Next
	}

	flag := 0
	result := (*ListNode)(nil)
	len1, len2 := len(s1), len(s2)
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
