package leetcode

import . "github.com/lijinglin2019/algorithm-go/leetcode"

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	childrenMaxDepth := 0
	for _, n := range root.Children {
		depth := maxDepth(n)
		if childrenMaxDepth < depth {
			childrenMaxDepth = depth
		}
	}

	return childrenMaxDepth + 1
}
