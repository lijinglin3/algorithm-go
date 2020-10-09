package leetcode

import "math"

func verifyPostorder(postorder []int) bool {
	length := len(postorder)
	stack := make([]int, 0)
	root := math.MaxInt64
	for i := length - 1; i >= 0; i-- {
		if postorder[i] > root {
			return false
		}
		for len(stack) != 0 && postorder[i] < stack[len(stack)-1] {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, postorder[i])
	}
	return true
}
