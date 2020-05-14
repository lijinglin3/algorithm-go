package leetcode

import . "github.com/lijinglin2019/algorithm-go/leetcode"

func hasPathSum(root *TreeNode, sum int) bool {
	if nil == root {
		return false
	}

	sumMap := map[*TreeNode]int{root: root.Val}
	stack := []*TreeNode{root}

	for len(stack) != 0 {
		i := len(stack) - 1
		n := stack[i]
		stack = stack[:i]
		s := sumMap[n]

		if n.Left == nil && n.Right == nil {
			if sum == s {
				return true
			}
			continue
		}
		if n.Left != nil {
			stack = append(stack, n.Left)
			sumMap[n.Left] = s + n.Left.Val
		}
		if n.Right != nil {
			stack = append(stack, n.Right)
			sumMap[n.Right] = s + n.Right.Val
		}
	}

	return false
}
