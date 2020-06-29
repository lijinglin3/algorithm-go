package leetcode

import . "github.com/lijinglin2019/algorithm-go/leetcode"

var res [][]int

func pathSum(root *TreeNode, sum int) [][]int {
	res = make([][]int, 0)
	dfs(root, sum, []int{})
	return res
}

func dfs(root *TreeNode, sum int, path []int) {
	if root == nil {
		return
	}

	path = append(path, root.Val)
	if root.Left == nil && root.Right == nil && sum == root.Val {
		r := make([]int, len(path))
		copy(r, path)
		res = append(res, r)
	}
	dfs(root.Left, sum-root.Val, path)
	dfs(root.Right, sum-root.Val, path)
}
