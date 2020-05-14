package leetcode

import . "github.com/lijinglin2019/algorithm-go/leetcode"

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	n := root
	for n != nil {
		switch {
		case n.Val == val:
			return n
		case n.Val > val:
			n = n.Left
		case n.Val < val:
			n = n.Right
		}
	}
	return nil
}
