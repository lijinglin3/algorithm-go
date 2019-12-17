package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	var ln *ListNode
	for head != nil {
		ln = &ListNode{Val: head.Val, Next: ln}
		head = head.Next
	}

	return ln
}
