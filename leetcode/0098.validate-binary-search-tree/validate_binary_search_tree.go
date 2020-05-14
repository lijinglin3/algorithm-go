package leetcode

import (
	"math"

	. "github.com/lijinglin2019/algorithm-go/leetcode/common"
)

func isValidBST(root *TreeNode) bool {
	min := math.MinInt64
	stack := make([]*TreeNode, 0)
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= min {
			return false
		}
		min = root.Val
		root = root.Right
	}
	return true
}
