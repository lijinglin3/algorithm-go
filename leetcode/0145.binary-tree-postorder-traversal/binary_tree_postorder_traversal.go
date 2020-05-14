package leetcode

import . "github.com/lijinglin2019/algorithm-go/leetcode"

func postorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return nil
	}

	result = append(result, postorderTraversal(root.Left)...)
	result = append(result, postorderTraversal(root.Right)...)
	result = append(result, root.Val)

	return result
}

func postorderTraversalByStack(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var result []int
	var pre *TreeNode
	stack := []*TreeNode{root}

	for len(stack) != 0 {
		i := len(stack) - 1
		n := stack[i]

		if (n.Right == nil && n.Left == nil) || (pre != nil && (pre == n.Left || pre == n.Right)) {
			result = append(result, n.Val)
			stack = stack[:i]
			pre = n
		} else {
			if n.Right != nil {
				stack = append(stack, n.Right)
			}
			if n.Left != nil {
				stack = append(stack, n.Left)
			}
		}
	}

	return result
}
