package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	var ln *ListNode
	for head != nil {
		cur := head
		head = head.Next
		cur.Next = ln
		ln = cur
	}

	return ln
}

func ReverseListByRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	ln := ReverseListByRecursion(head.Next)
	head.Next.Next = head
	head.Next = nil

	return ln
}
