package leetcode

import . "github.com/lijinglin2019/algorithm-go/leetcode/common"

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	p := head
	var l *ListNode
	for {
		p = p.Next
		if p == nil {
			p = head.Next
			break
		}

		tmp := head
		head, p = head.Next, p.Next
		tmp.Next = l
		l = tmp
		if p == nil {
			p = head
			break
		}
	}

	for l != nil && p != nil {
		if l.Val != p.Val {
			return false
		}
		l, p = l.Next, p.Next
	}

	return true
}
