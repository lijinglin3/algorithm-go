package leetcode

import . "github.com/lijinglin2019/algorithm-go/leetcode"

func isSymmetric(root *TreeNode) bool {
	queue := []*TreeNode{root, root}
	for len(queue) > 0 {
		n1 := queue[len(queue)-1]
		n2 := queue[len(queue)-2]
		queue = queue[:len(queue)-2]

		if n1 == nil && n2 == nil {
			continue
		}
		if n1 == nil || n2 == nil {
			return false
		}
		if n1.Val != n2.Val {
			return false
		}
		queue = append(queue, n1.Left, n2.Right, n1.Right, n2.Left)
	}
	return true
}

func isSymmetricByRecursive(root *TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	return (root1.Val == root2.Val) &&
		isMirror(root1.Left, root2.Right) &&
		isMirror(root1.Right, root2.Left)
}
