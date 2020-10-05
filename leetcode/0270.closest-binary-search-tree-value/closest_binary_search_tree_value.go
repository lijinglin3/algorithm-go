package leetcode

import (
	"math"

	. "github.com/lijinglin3/algorithm-go/leetcode"
)

func closestValue(root *TreeNode, target float64) int {
	delta, result := math.Abs(float64(root.Val)-target), root.Val
	for root != nil {
		tmp := math.Abs(float64(root.Val) - target)
		if delta >= tmp {
			delta, result = tmp, root.Val
		}
		if float64(root.Val) > target {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return result
}
