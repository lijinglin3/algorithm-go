package leetcode

import . "github.com/lijinglin2019/algorithm-go/leetcode"

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	tmp := make(map[*TreeNode]int)
	stack, sum := []*TreeNode{root}, 0
	for len(stack) != 0 {
		length := len(stack) - 1
		n := stack[length]
		stack = stack[:length]
		if v, ok := tmp[n]; ok && n.Left == nil && n.Right == nil {
			sum += v
		}
		if n.Left != nil {
			stack = append(stack, n.Left)
			tmp[n.Left] = n.Left.Val
		}
		if n.Right != nil {
			stack = append(stack, n.Right)
		}
	}
	return sum
}
