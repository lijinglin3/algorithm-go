package leetcode

import . "github.com/lijinglin2019/algorithm-go/leetcode"

var count map[int]int

func pathSum(root *TreeNode, sum int) int {
	count = map[int]int{0: 1}
	return dfs(root, sum, 0)
}

func dfs(root *TreeNode, sum int, prefix int) int {
	if root == nil {
		return 0
	}
	prefix += root.Val
	res := count[prefix-sum]
	count[prefix]++
	res += dfs(root.Left, sum, prefix) + dfs(root.Right, sum, prefix)
	count[prefix]--
	return res
}
