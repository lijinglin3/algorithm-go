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

//func verifyPostorder(postorder []int) bool {
//	length := len(postorder)
//	if length == 0 {
//		return true
//	}
//	f := length - 1
//	for i := 0; i < length; i++ {
//		if postorder[i] >= postorder[length-1] {
//			f = i
//		}
//		if i > f && postorder[i] < postorder[length-1] {
//			return false
//		}
//	}
//	return verifyPostorder(postorder[:f]) && verifyPostorder(postorder[f:length-1])
//}
