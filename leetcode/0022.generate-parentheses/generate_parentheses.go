package leetcode

var output []string

func generateParenthesis(n int) []string {
	output = make([]string, 0)
	iteration(0, 0, n, "")
	return output
}

func iteration(left int, right int, max int, s string) {
	if left == right && left == max {
		output = append(output, s)
		return
	}

	if left < max {
		iteration(left+1, right, max, s+"(")
	}
	if right < left {
		iteration(left, right+1, max, s+")")
	}
}
