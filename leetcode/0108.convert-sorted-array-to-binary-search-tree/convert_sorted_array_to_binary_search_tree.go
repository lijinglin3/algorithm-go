package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SortedArrayToBST(nums []int) *TreeNode {
	length := len(nums)
	if length == 0 {
		return nil
	}

	mid := length / 2
	left := nums[:mid]
	right := nums[mid+1:]
	return &TreeNode{Val: nums[mid], Left: SortedArrayToBST(left), Right: SortedArrayToBST(right)}
}
