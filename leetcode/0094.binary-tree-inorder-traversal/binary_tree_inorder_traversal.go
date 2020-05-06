package leetcode

import . "github.com/lijinglin2019/algorithm-go/leetcode/common"

func InorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return nil
	}

	result = append(result, InorderTraversal(root.Left)...)
	result = append(result, root.Val)
	result = append(result, InorderTraversal(root.Right)...)

	return result
}
