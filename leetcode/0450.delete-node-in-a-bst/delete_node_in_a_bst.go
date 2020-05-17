package leetcode

import . "github.com/lijinglin2019/algorithm-go/leetcode"

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	switch {
	case root.Val > key:
		root.Left = deleteNode(root.Left, key)
	case root.Val < key:
		root.Right = deleteNode(root.Right, key)
	default:
		switch {
		case root.Left == nil && root.Right == nil:
			root = nil
		case root.Right != nil:
			root.Val = left(root.Right)
			root.Right = deleteNode(root.Right, root.Val)
		default:
			root.Val = right(root.Left)
			root.Left = deleteNode(root.Left, root.Val)
		}
	}
	return root
}

func left(root *TreeNode) int {
	for root.Left != nil {
		root = root.Left
	}
	return root.Val
}

func right(root *TreeNode) int {
	for root.Right != nil {
		root = root.Right
	}
	return root.Val
}
