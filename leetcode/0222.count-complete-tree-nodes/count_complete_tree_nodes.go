package leetcode

import (
	. "github.com/lijinglin3/algorithm-go/leetcode"
)

func countNodes(root *TreeNode) int {
	n := root
	depth := 0
	for n != nil {
		depth++
		n = n.Left
	}

	if depth <= 1 {
		return depth
	}

	l, r := 1, 1<<(depth-1)
	for l <= r {
		mid := (l + r) >> 1
		if check(mid, depth, root) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return 1<<(depth-1) - 1 + r
}

func check(index, depth int, root *TreeNode) bool {
	l, r := 1, 1<<(depth-1)
	for i := 0; i < depth-1; i++ {
		pivot := (l + r) >> 1
		if index <= pivot {
			r = pivot
			root = root.Left
		} else {
			l = pivot + 1
			root = root.Right
		}
	}
	return root == nil
}
