package leetcode

import . "github.com/lijinglin3/algorithm-go/leetcode"

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for {
		if fast == nil || fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}

	for slow = head; slow != fast; slow = slow.Next {
		fast = fast.Next
	}

	return slow
}
