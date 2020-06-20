package leetcode

func dailyTemperatures(T []int) []int {
	stack, result := make([]int, 0), make([]int, len(T))
	for i := 0; i < len(T); i++ {
		for len(stack) > 0 && T[stack[len(stack)-1]] < T[i] {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[index] = i - index
		}
		stack = append(stack, i)
	}
	return result
}
