package leetcode

func validateStackSequences(pushed []int, popped []int) bool {
	stack, index := make([]int, 0, len(pushed)), 0
	for i := 0; i < len(popped); i++ {
		stack = append(stack, pushed[i])
		for len(stack) > 0 && stack[len(stack)-1] == popped[index] {
			stack = stack[:len(stack)-1]
			index++
		}
	}
	return len(stack) == 0
}
