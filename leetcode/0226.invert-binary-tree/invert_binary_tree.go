package leetcode

import . "github.com/lijinglin3/algorithm-go/leetcode"

func invertTree(root *TreeNode) *TreeNode {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}

	left := invertTree(root.Left)
	right := invertTree(root.Right)

	root.Left, root.Right = right, left
	return root
}
