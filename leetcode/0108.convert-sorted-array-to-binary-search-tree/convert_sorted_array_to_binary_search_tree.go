package leetcode

import . "github.com/lijinglin2019/algorithm-go/leetcode"

func sortedArrayToBST(nums []int) *TreeNode {
	length := len(nums)
	if length == 0 {
		return nil
	}

	mid := length / 2
	left := nums[:mid]
	right := nums[mid+1:]
	return &TreeNode{Val: nums[mid], Left: sortedArrayToBST(left), Right: sortedArrayToBST(right)}
}
