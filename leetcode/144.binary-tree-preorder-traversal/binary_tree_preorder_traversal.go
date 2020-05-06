package leetcode

import . "github.com/lijinglin2019/algorithm-go/leetcode/common"

func PreorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return nil
	}

	result = append(result, root.Val)
	result = append(result, PreorderTraversal(root.Left)...)
	result = append(result, PreorderTraversal(root.Right)...)

	return result
}
