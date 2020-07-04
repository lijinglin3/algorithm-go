package leetcode

import . "github.com/lijinglin2019/algorithm-go/leetcode"

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	_, b := checkBalanced(root)
	return b
}

func checkBalanced(root *TreeNode) (int, bool) {
	if root.Left == nil && root.Right == nil {
		return 1, true
	}

	left := 0
	if root.Left != nil {
		if l, balanced := checkBalanced(root.Left); balanced {
			left = l
		} else {
			return 0, false
		}
	}

	right := 0
	if root.Right != nil {
		if r, balanced := checkBalanced(root.Right); balanced {
			right = r
		} else {
			return 0, false
		}
	}

	max, diff := 0, 0
	if left > right {
		max = left
		diff = left - right
	} else {
		max = right
		diff = right - left
	}

	if diff > 1 {
		return 0, false
	}
	return max + 1, true
}
