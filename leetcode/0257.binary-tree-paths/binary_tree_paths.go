package leetcode

import (
	"strconv"

	. "github.com/lijinglin2019/algorithm-go/leetcode"
)

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}

	result := make([]string, 0)
	stack, paths := []*TreeNode{root}, []string{strconv.FormatInt(int64(root.Val), 10)}
	for len(stack) != 0 {
		length := len(stack) - 1
		n, p := stack[length], paths[length]
		stack, paths = stack[:length], paths[:length]
		if n.Left == nil && n.Right == nil {
			result = append(result, p)
		}
		if n.Left != nil {
			stack = append(stack, n.Left)
			paths = append(paths, p+"->"+strconv.FormatInt(int64(n.Left.Val), 10))
		}
		if n.Right != nil {
			stack = append(stack, n.Right)
			paths = append(paths, p+"->"+strconv.FormatInt(int64(n.Right.Val), 10))
		}
	}
	return result
}
