package leetcode

func generatePossibleNextMoves(s string) []string {
	if len(s) < 2 {
		return nil
	}
	result := make([]string, 0)
	for i := 1; i < len(s); {
		if s[i] == '+' {
			if s[i-1] == '+' {
				result = append(result, s[:i-1]+"--"+s[i+1:])
			}
			i++
		} else {
			i += 2
		}
	}
	return result
}
