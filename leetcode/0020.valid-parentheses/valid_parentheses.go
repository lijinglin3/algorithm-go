package leetcode

func isValid(s string) bool {
	stack, index := make([]rune, len(s)), 0
	for _, v := range s {
		if index != 0 && ((stack[index-1] == '{' && v == '}') || (stack[index-1] == '[' && v == ']') || (stack[index-1] == '(' && v == ')')) {
			index--
		} else {
			stack[index] = v
			index++
		}
	}
	return index == 0
}
