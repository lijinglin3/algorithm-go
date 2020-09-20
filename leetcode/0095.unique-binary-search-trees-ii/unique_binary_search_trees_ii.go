package leetcode

import . "github.com/lijinglin2019/algorithm-go/leetcode"

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return generate(1, n)
}

func generate(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}

	var results []*TreeNode
	for i := start; i <= end; i++ {
		leftTrees := generate(start, i-1)
		rightTrees := generate(i+1, end)
		for _, left := range leftTrees {
			for _, right := range rightTrees {
				currTree := &TreeNode{Val: i}
				currTree.Left = left
				currTree.Right = right
				results = append(results, currTree)
			}
		}
	}
	return results
}
