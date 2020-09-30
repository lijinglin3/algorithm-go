package leetcode

import . "github.com/lijinglin2019/algorithm-go/leetcode"

func findMode(root *TreeNode) (answer []int) {
	var base, count, maxCount int

	update := func(x int) {
		if x == base {
			count++
		} else {
			base, count = x, 1
		}
		if count == maxCount {
			answer = append(answer, base)
		} else if count > maxCount {
			maxCount = count
			answer = []int{base}
		}
	}

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		update(node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return
}
