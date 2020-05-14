package leetcode

import . "github.com/lijinglin2019/algorithm-go/leetcode/common"

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
