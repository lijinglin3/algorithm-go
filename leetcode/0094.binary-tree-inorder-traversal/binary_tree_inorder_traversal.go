package leetcode

import . "github.com/lijinglin3/algorithm-go/leetcode"

func inorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return nil
	}

	result = append(result, inorderTraversal(root.Left)...)
	result = append(result, root.Val)
	result = append(result, inorderTraversal(root.Right)...)

	return result
}

func inorderTraversalByStack(root *TreeNode) []int {
	var result []int
	stack := make([]*TreeNode, 0)
	n := root

	for len(stack) != 0 || n != nil {
		if n != nil {
			stack = append(stack, n)
			n = n.Left
		} else {
			i := len(stack) - 1
			result = append(result, stack[i].Val)
			n = stack[i].Right
			stack = stack[:i]
		}
	}

	return result
}
