package leetcode

import . "github.com/lijinglin3/algorithm-go/leetcode"

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	n := &TreeNode{Val: nums[mid]}
	n.Left = sortedArrayToBST(nums[:mid])
	n.Right = sortedArrayToBST(nums[mid+1:])
	return n
}
