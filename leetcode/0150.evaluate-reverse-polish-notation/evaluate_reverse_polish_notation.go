package leetcode

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int64, 0)
	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "+":
			i1, i2 := stack[len(stack)-2], stack[len(stack)-1]
			stack[len(stack)-2] = i1 + i2
			stack = stack[:len(stack)-1]
		case "-":
			i1, i2 := stack[len(stack)-2], stack[len(stack)-1]
			stack[len(stack)-2] = i1 - i2
			stack = stack[:len(stack)-1]
		case "*":
			i1, i2 := stack[len(stack)-2], stack[len(stack)-1]
			stack[len(stack)-2] = i1 * i2
			stack = stack[:len(stack)-1]
		case "/":
			i1, i2 := stack[len(stack)-2], stack[len(stack)-1]
			stack[len(stack)-2] = i1 / i2
			stack = stack[:len(stack)-1]
		default:
			v, _ := strconv.ParseInt(tokens[i], 10, 64)
			stack = append(stack, v)
		}
	}
	return int(stack[0])
}
