package leetcode

import . "github.com/lijinglin2019/algorithm-go/leetcode"

func kthLargest(root *TreeNode, k int) int {
	result := make([]int, 0, k)
	traverse(root, &result, k)
	return result[len(result)-1]
}

func traverse(root *TreeNode, count *[]int, k int) {
	if root == nil || len(*count) >= k {
		return
	}
	traverse(root.Right, count, k)
	if len(*count) >= k {
		return
	}
	*count = append(*count, root.Val)
	traverse(root.Left, count, k)
}
