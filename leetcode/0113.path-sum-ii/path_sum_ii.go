package leetcode

import . "github.com/lijinglin2019/algorithm-go/leetcode"

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}

	result := make([][]int, 0)
	stack := []*TreeNode{root}
	paths := [][]int{{root.Val}}
	sums := []int{root.Val}

	for len(stack) != 0 {
		length := len(stack) - 1
		n, p, s := stack[length], paths[length], sums[length]
		stack, paths, sums = stack[:length], paths[:length], sums[:length]
		if n.Left == nil && n.Right == nil && s == sum {
			result = append(result, p)
		}

		if n.Left != nil {
			stack = append(stack, n.Left)
			tmp := make([]int, len(p))
			copy(tmp, p)
			paths = append(paths, append(tmp, n.Left.Val))
			sums = append(sums, s+n.Left.Val)
		}
		if n.Right != nil {
			stack = append(stack, n.Right)
			tmp := make([]int, len(p))
			copy(tmp, p)
			paths = append(paths, append(tmp, n.Right.Val))
			sums = append(sums, s+n.Right.Val)
		}
	}
	return result
}

var res = make([][]int, 0)

func pathSum2(root *TreeNode, sum int) [][]int {
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
