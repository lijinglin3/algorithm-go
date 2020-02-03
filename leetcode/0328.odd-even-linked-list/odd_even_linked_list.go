package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func OddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}

	i := 3
	odd, cur := head, head.Next
	for cur.Next != nil {
		if i%2 != 0 {
			tmp := cur.Next
			cur.Next = cur.Next.Next
			tmp.Next = odd.Next
			odd.Next = tmp
			odd = odd.Next
		} else {
			cur = cur.Next
		}
		i++
	}

	return head
}
