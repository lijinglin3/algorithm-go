package leetcode

import . "github.com/lijinglin2019/algorithm-go/leetcode"

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	n := root
	for {
		switch {
		case val > n.Val:
			if n.Right != nil {
				n = n.Right
			} else {
				n.Right = &TreeNode{Val: val}
				return root
			}
		case val < n.Val:
			if n.Left != nil {
				n = n.Left
			} else {
				n.Left = &TreeNode{Val: val}
				return root
			}
		default:
			return root
		}
	}
}
