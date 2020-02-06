package leetcode

func isValid(s string) bool {
	stack := make([]rune, 0)
	for _, v := range s {
		if len(stack) == 0 {
			stack = append(stack, v)
			continue
		}
		head := stack[len(stack)-1]
		switch {
		case string(head) == "{" && string(v) == "}":
			stack = stack[:len(stack)-1]
		case string(head) == "[" && string(v) == "]":
			stack = stack[:len(stack)-1]
		case string(head) == "(" && string(v) == ")":
			stack = stack[:len(stack)-1]
		default:
			stack = append(stack, v)
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}
