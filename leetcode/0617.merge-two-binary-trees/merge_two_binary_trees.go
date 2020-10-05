package leetcode

import . "github.com/lijinglin2019/algorithm-go/leetcode"

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return copyTree(t2)
	}
	if t2 == nil {
		return copyTree(t1)
	}

	newT := &TreeNode{Val: t1.Val + t2.Val}
	newT.Right = mergeTrees(t1.Right, t2.Right)
	newT.Left = mergeTrees(t1.Left, t2.Left)
	return newT
}

func copyTree(t *TreeNode) *TreeNode {
	if t == nil {
		return nil
	}

	newT := &TreeNode{Val: t.Val}
	newT.Right = copyTree(t.Right)
	newT.Left = copyTree(t.Left)
	return newT
}
