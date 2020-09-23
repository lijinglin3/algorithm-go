package leetcode

func permutation(s string) []string {
	if len(s) == 0 {
		return []string{""}
	}
	result := make([]string, 0)
	dict := make(map[string]bool)
	str := []byte(s)
	var backtrack func(index int)
	backtrack = func(index int) {
		if index == len(str) {
			ss := string(str)
			if !dict[ss] {
				dict[ss] = true
				result = append(result, ss)
			}
			return
		}
		for i := index; i < len(str); i++ {
			tmp := str[index]
			str[index] = str[i]
			str[i] = tmp
			backtrack(index + 1)
			str[i] = str[index]
			str[index] = tmp
		}
	}
	backtrack(0)
	return result
}
