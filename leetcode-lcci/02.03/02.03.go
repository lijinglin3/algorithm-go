package leetcode

import . "github.com/lijinglin3/algorithm-go/leetcode"

func deleteNode(node *ListNode) {
	//node.Val = node.Next.Val
	//node.Next = node.Next.Next
	*node = *node.Next
}
