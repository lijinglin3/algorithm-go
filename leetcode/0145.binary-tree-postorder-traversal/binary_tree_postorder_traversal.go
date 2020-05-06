package leetcode

import . "github.com/lijinglin2019/algorithm-go/leetcode/common"

func PostorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return nil
	}

	result = append(result, PostorderTraversal(root.Left)...)
	result = append(result, PostorderTraversal(root.Right)...)
	result = append(result, root.Val)

	return result
}
