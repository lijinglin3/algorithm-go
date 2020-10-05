package leetcode

import . "github.com/lijinglin3/algorithm-go/leetcode"

func kthSmallest(root *TreeNode, k int) int {
	stack := make([]*TreeNode, 0)
	n := root

	count := 0
	for len(stack) != 0 || n != nil {
		if n != nil {
			stack = append(stack, n)
			n = n.Left
		} else {
			i := len(stack) - 1
			count++
			if count == k {
				return stack[i].Val
			}
			n = stack[i].Right
			stack = stack[:i]
		}
	}
	return -1
}
