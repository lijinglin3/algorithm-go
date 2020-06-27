package leetcode

import . "github.com/lijinglin2019/algorithm-go/leetcode"

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left, right := mirrorTree(root.Left), mirrorTree(root.Right)
	root.Left, root.Right = right, left
	return root
}
