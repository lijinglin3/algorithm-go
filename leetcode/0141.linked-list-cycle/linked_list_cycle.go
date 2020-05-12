package leetcode

// ListNode definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func HasCycle(head *ListNode) bool {
	p1, p2 := head, head

	for {
		if p2 == nil {
			return false
		}

		p1, p2 = p1.Next, p2.Next
		if p2 == nil {
			return false
		}

		p2 = p2.Next
		if p1 == p2 {
			return true
		}
	}
}
