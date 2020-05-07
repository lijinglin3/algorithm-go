package leetcode

import . "github.com/lijinglin2019/algorithm-go/leetcode/common"

func PreorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return nil
	}

	result = append(result, root.Val)
	result = append(result, PreorderTraversal(root.Left)...)
	result = append(result, PreorderTraversal(root.Right)...)

	return result
}

func PreorderTraversalByStack(root *TreeNode) []int {
	var result []int
	stack := make([]*TreeNode, 0)
	n := root

	for len(stack) != 0 || n != nil {
		if n != nil {
			result = append(result, n.Val)
			stack = append(stack, n)
			n = n.Left
		} else {
			i := len(stack) - 1
			n = stack[i].Right
			stack = stack[:i]
		}
	}

	return result
}
